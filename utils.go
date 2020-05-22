package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type SSR struct {
	server        string
	serverPort    int
	localAddress  string
	localPort     int
	timeout       int
	workers       int
	password      string
	method        string // TODO
	obfs          string // TODO
	obfsParam     string
	protocol      string // TODO
	protocolParam string
}

type SS struct {
	server       string
	serverPort   int
	localAddress string
	localPort    int
	timeout      int
	workers      int
	password     string
	method       string // TODO
	plugin       string
}

func safeDecode(raw []byte) []byte {
	var ret []byte
	safeLen := len(raw) - len(raw)%4
	safe, rest := raw[0:safeLen], raw[safeLen:]
	decodedSafe, err := base64.StdEncoding.DecodeString(string(safe))
	if err != nil {
		fmt.Println("[warning]", err)
	}

	var decodeMap [256]uint8
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	encoder := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
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

func decodeLink(link string) *SSR {
	fmt.Printf("[link] %s\n", link)
	return &SSR{}
}

func getSSR(url string, cache bool) ([]*SSR, error) {
	var raw []byte
	CacheFile := os.TempDir() + "/" + base64.RawStdEncoding.EncodeToString([]byte(url))
	if _, err := os.Stat(CacheFile); cache && !os.IsNotExist(err) {
		fmt.Println("Loading from cache...")
		raw, _ = ioutil.ReadFile(CacheFile)
	} else {
		fmt.Println("Loading from web...")
		res, err := (&http.Client{Timeout: 10 * time.Second}).Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		raw, _ = ioutil.ReadAll(res.Body)
		ioutil.WriteFile(CacheFile, raw, 0755)
	}
	links := strings.Split(strings.TrimSpace(string(safeDecode(raw))), "\n")
	var res []*SSR
	for _, link := range links {
		decodedLink := regexp.MustCompile(`[^/]+$`).ReplaceAllStringFunc(link, func(s string) string {
			var decodedHalf []string
			for _, half := range strings.Split(s, "_") {
				decodedHalf = append(decodedHalf, string(safeDecode([]byte(half))))
			}
			return strings.Join(decodedHalf, "?")
		})
		res = append(res, decodeLink(decodedLink))
	}

	return res, nil
}
