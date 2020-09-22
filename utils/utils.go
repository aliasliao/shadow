package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aliasliao/shadow/model"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

type ShadowsocksR struct {
	server        string
	serverPort    uint32
	localAddress  string
	localPort     uint32
	timeout       uint32
	workers       uint32
	password      string
	method        string // TODO enum
	obfs          string // TODO enum
	obfsParam     string
	protocol      string // TODO enum
	protocolParam string
	group         string
	remarks       string
}
type Shadowsocks struct {
	server       string
	serverPort   uint32
	localAddress string
	localPort    uint32
	timeout      uint32
	workers      uint32
	password     string
	method       string // TODO enum
	plugin       string
	group        string
	remarks      string
}

type options struct {
	loglevel string
}

func safeDecode(raw []byte) []byte {
	var ret []byte
	safeLen := len(raw) - len(raw)%4
	safe, rest := raw[0:safeLen], raw[safeLen:]
	decodedSafe, err := (func() ([]byte, error) {
		if strings.ContainsAny(string(safe), "+/") {
			return base64.StdEncoding.DecodeString(string(safe))
		}
		return base64.URLEncoding.DecodeString(string(safe))
	})()
	if err != nil {
		log.Println("[warning](base64 error)", err)
		log.Panic("raw---->", string(raw))
	}

	var decodeMap [256]uint8
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	encoder := (func() string {
		if strings.ContainsAny(string(safe), "+/") {
			return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
		} else {
			return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
		}
	})()
	for i := 0; i < len(encoder); i++ {
		decodeMap[encoder[i]] = uint8(i)
	}

	var bin uint32 = 0
	for _, b := range rest {
		bin = bin<<6 | uint32(decodeMap[b])
	}
	bin = bin >> (len(rest) * 6 % 8)
	var decodedRest []byte
	for bin > 0 {
		decodedRest = append([]byte{uint8(bin & 0xFF)}, decodedRest...)
		bin = bin >> 8
	}

	ret = append(ret, decodedSafe...)
	ret = append(ret, decodedRest...)
	return ret
}

func safeDecodeStr(raw string) string {
	return string(safeDecode([]byte(raw)))
}

func decodeSSRLink(link string) (*ShadowsocksR, error) {
	decodedLink := regexp.MustCompile(`[^/]+$`).ReplaceAllStringFunc(link, func(s string) string {
		var decodedHalf []string
		for _, half := range strings.Split(s, "_") {
			decodedHalf = append(decodedHalf, safeDecodeStr(half))
		}
		return strings.Join(decodedHalf, "?")
	})
	re := regexp.MustCompile(`ssr://(?P<server>[^:]+):(?P<serverPort>[^:]+):(?P<protocol>[^:]+):(?P<method>[^:]+):(?P<obfs>[^:]+):(?P<password>[^:]+?)/\?(?P<queries>[\S]+)$`)
	keys := re.SubexpNames()
	values := re.FindStringSubmatch(decodedLink)
	if values == nil {
		return nil, errors.New(fmt.Sprintf("unsupported format (%s)", decodedLink))
	}
	captured := map[string]string{}
	for index, key := range keys {
		captured[key] = values[index]
	}
	queries, err := url.ParseQuery(captured["queries"])
	if err != nil {
		return nil, err
	}
	for key, value := range queries {
		captured[key] = value[0]
	}
	serverPort, err := strconv.ParseUint(captured["serverPort"], 10, 16)
	if err != nil {
		return nil, err
	}

	return &ShadowsocksR{
		server:        captured["server"],
		serverPort:    uint32(serverPort),
		localAddress:  "",
		localPort:     0,
		timeout:       0,
		workers:       0,
		password:      safeDecodeStr(captured["password"]),
		method:        captured["method"],
		obfs:          captured["obfs"],
		obfsParam:     safeDecodeStr(captured["obfsparam"]),
		protocol:      captured["protocol"],
		protocolParam: safeDecodeStr(captured["protoparam"]),
		group:         safeDecodeStr(captured["group"]),
		remarks:       safeDecodeStr(captured["remarks"]),
	}, nil
}

func decodeSSLink(link string) (*Shadowsocks, error) {
	re := regexp.MustCompile(`^ss://(?P<methodAndPassword>\S+)@(?P<server>\S+):(?P<serverPort>\S+)/\?group=(?P<group>\S+)#(?P<remarks>\S+)$`)
	keys := re.SubexpNames()
	values := re.FindStringSubmatch(link)
	if values == nil {
		return nil, errors.New(fmt.Sprintf("unsupported format (%s)", link))
	}
	captured := map[string]string{}
	for index, key := range keys {
		captured[key] = values[index]
	}

	serverPort, err := strconv.ParseUint(captured["serverPort"], 10, 16)
	if err != nil {
		return nil, err
	}
	methodAndPassword := strings.Split(safeDecodeStr(captured["methodAndPassword"]), ":")
	remarks, err := url.QueryUnescape(captured["remarks"])
	if err != nil {
		return nil, err
	}

	return &Shadowsocks{
		server:       captured["server"],
		serverPort:   uint32(serverPort),
		localAddress: "",
		localPort:    0,
		timeout:      0,
		workers:      0,
		password:     methodAndPassword[1],
		method:       methodAndPassword[0],
		plugin:       "",
		group:        safeDecodeStr(captured["group"]),
		remarks:      remarks,
	}, nil
}

// fetchHTTPContent dials https for remote content
func fetchHTTPContent(target string) ([]byte, error) {
	res, err := (&http.Client{Timeout: 30 * time.Second}).Get(target)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	return content, nil
}

func loadRaw(url string, cache bool) ([]byte, error) {
	if cache {
		CacheFile := os.TempDir() + "/" + base64.RawURLEncoding.EncodeToString([]byte(url))
		if _, err := os.Stat(CacheFile); !os.IsNotExist(err) {
			log.Println("Loading from cache...")
			return ioutil.ReadFile(CacheFile)
		} else {
			log.Println("Loading from web...")
			raw, err := fetchHTTPContent(url)
			if err != nil {
				return nil, err
			}
			ioutil.WriteFile(CacheFile, raw, 0755)
			log.Println("File saved to", CacheFile)
			return raw, nil
		}
	}
	log.Println("Loading from web...")
	return fetchHTTPContent(url)
}

func parseSSR(url string, cache bool) ([]*ShadowsocksR, error) {
	raw, err := loadRaw(url, cache)
	if err != nil {
		return nil, err
	}
	links := strings.Split(strings.TrimSpace(string(safeDecode(raw))), "\n")
	var list []*ShadowsocksR
	for _, link := range links {
		ssr, err := decodeSSRLink(link)
		if err != nil {
			log.Println("[warning](bad link skipped)", err)
			continue
		}
		list = append(list, ssr)
	}

	return list, nil
}

func parseSSD(url string, cache bool) ([]*Shadowsocks, error) {
	type Server struct {
		Id      uint32
		Server  string
		Ratio   uint32
		Remarks string
	}
	type SSD struct {
		Airport       string
		Port          uint32
		Encryption    string
		Password      string
		Traffic_used  float32
		Traffic_total float32
		Expiry        string
		Url           string
		Servers       []Server
	}

	raw, err := loadRaw(url, cache)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`ssd://([a-zA-Z0-9+/]+)`) // encodeStd
	js := safeDecodeStr(re.FindStringSubmatch(string(raw))[1])
	var ssd *SSD
	if err := json.Unmarshal([]byte(js), &ssd); err != nil {
		return nil, err
	}

	var list []*Shadowsocks
	for _, server := range ssd.Servers {
		ss := &Shadowsocks{
			server:       server.Server,
			serverPort:   ssd.Port,
			localAddress: "",
			localPort:    0,
			timeout:      0,
			workers:      0,
			password:     ssd.Password,
			method:       ssd.Encryption,
			plugin:       "",
			remarks:      server.Remarks,
		}
		list = append(list, ss)
	}

	return list, nil
}

func parseSS(url string, cache bool) ([]*Shadowsocks, error) {
	raw, err := loadRaw(url, cache)
	if err != nil {
		return nil, err
	}
	links := strings.Split(strings.TrimSpace(string(safeDecode(raw))), "\n")
	var list []*Shadowsocks
	for _, link := range links {
		ss, err := decodeSSLink(link)
		if err != nil {
			log.Println("[warning](bad link skipped", err)
			continue
		}
		list = append(list, ss)
	}

	return list, nil
}

func ssToConfig(sss []*Shadowsocks, options *options) *model.Config {
	var servers []*model.ShadowsocksOutboundConfigurationObject_ServerObject
	for _, ss := range sss {
		servers = append(servers, &model.ShadowsocksOutboundConfigurationObject_ServerObject{
			Email:    "",
			Address:  ss.server,
			Port:     ss.serverPort,
			Method:   ss.method,
			Password: ss.password,
			Ota:      false,
			Level:    0,
		})
	}

	config := &model.Config{
		Inbounds: []*model.InboundObject{{
			Port:     1081,
			Listen:   "127.0.0.1",
			Protocol: "socks",
			Settings: nil,
			Sniffing: &model.InboundObject_SniffingObject{
				Enabled:      true,
				DestOverride: []string{"http", "tls"},
			},
		}, {
			Tag:      "api-in",
			Protocol: "dokodemo-door",
			Port:     8080,
			Listen:   "127.0.0.1",
			Settings: (func() *any.Any {
				ret, err := ptypes.MarshalAny(&model.DokodemoInboundConfigurationObject{
					Address: "127.0.0.1",
				})
				if err != nil {
				}
				return ret
			})(),
		}},
		Outbounds: []*model.OutboundObject{{
			Protocol: "shadowsocks",
			Settings: (func() *any.Any {
				ret, err := ptypes.MarshalAny(&model.ShadowsocksOutboundConfigurationObject{
					Servers: servers,
				})
				if err != nil {
				}
				return ret
			})(),
		}},
		Api: &model.ApiObject{
			Tag: "api-out",
			Services: []string{
				"StatsService",
				"AppService",
			},
		},
		Routing: &model.RoutingObject{
			Rules: []*model.RoutingObject_RuleObject{{
				Type:        "field",
				InboundTag:  []string{"api-in"},
				OutboundTag: "api-out",
			}},
		},
		Log: &model.LogObject{
			Loglevel: options.loglevel,
		},
	}

	return config
}
