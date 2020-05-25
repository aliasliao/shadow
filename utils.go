package main

import (
	"encoding/base64"
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
)

type SSR struct {
	server        string
	serverPort    uint16
	localAddress  string
	localPort     uint16
	timeout       uint16
	workers       uint16
	password      string
	method        string // TODO enum
	obfs          string // TODO enum
	obfsParam     string
	protocol      string // TODO enum
	protocolParam string
	group         string
	remarks       string
}

type SS struct {
	server       string
	serverPort   uint16
	localAddress string
	localPort    uint16
	timeout      uint16
	workers      uint16
	password     string
	method       string // TODO enum
	plugin       string
}

func safeDecode(raw []byte) []byte {
	var ret []byte
	safeLen := len(raw) - len(raw)%4
	safe, rest := raw[0:safeLen], raw[safeLen:]
	decodedSafe, err := base64.URLEncoding.DecodeString(string(safe))
	if err != nil {
		log.Println("[warning](base64 error)", err)
		log.Println("raw---->", string(raw))
	}

	var decodeMap [256]uint8
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	encoder := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
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

func decodeLink(link string) (*SSR, error) {
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

	return &SSR{
		server:        captured["server"],
		serverPort:    uint16(serverPort),
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

func loadRaw(url string, cache bool) ([]byte, error) {
	var raw []byte
	CacheFile := os.TempDir() + "/" + base64.RawURLEncoding.EncodeToString([]byte(url))
	if _, err := os.Stat(CacheFile); cache && !os.IsNotExist(err) {
		log.Println("Loading from cache...")
		raw, _ = ioutil.ReadFile(CacheFile)
	} else {
		log.Println("Loading from web...")
		res, err := (&http.Client{Timeout: 20 * time.Second}).Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		raw, _ = ioutil.ReadAll(res.Body)
		ioutil.WriteFile(CacheFile, raw, 0755)
	}
	return raw, nil
}

func getSSR(url string, cache bool) ([]*SSR, error) {
	raw, err := loadRaw(url, cache)
	if err != nil {
		return nil, err
	}
	links := strings.Split(strings.TrimSpace(string(safeDecode(raw))), "\n")
	var res []*SSR
	for _, link := range links {
		ssr, err := decodeLink(link)
		if err != nil {
			log.Println("[warning](bad link skipped)", err)
			continue
		}
		res = append(res, ssr)
	}

	return res, nil
}

func getSSD(url string, cache bool) ([]*SS, error) {
	raw, err := loadRaw(url, cache)
	if err != nil {
		return nil, err
	}
	log.Println("raw", string(raw), "\n---\n", safeDecodeStr(string(raw)))
	return nil, nil
}
