package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	ret := make([]byte, 0)
	safeLen := len(raw) - len(raw)%4
	safe, rest := raw[0:safeLen], raw[safeLen:]
	decodedSafe, err := base64.StdEncoding.DecodeString(string(safe))
	if err != nil {
		fmt.Println("[warning]", err)
	}

	var decodeMap [256]uint32
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	encoder := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for i := 0; i < len(encoder); i++ {
		decodeMap[encoder[i]] = uint32(i)
	}

	var bin uint32 = 0
	for _, b := range rest {
		bin = bin<<6 | decodeMap[b]
	}
	// TODO
	ret = append(ret, decodedSafe...)
	return ret
}

func decodeLink(link string) *SSR {
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
	links := strings.Split(string(safeDecode(raw)), "\n")
	res := make([]*SSR, 0)
	for _, link := range links {
		res = append(res, decodeLink(link))
	}

	return res, nil
	//// ssr://120.232.150.53:65533:auth_aes128_md5:chacha20-ietf:tls1.2_ticket_auth:cXdlcnQ/?obfsparam=ZTIwMDgyNzUzOS5kb3dubG9hZC53aW5kb3dzdXBkYXRlLmNvbQ&protoparam=Mjc1Mzk6RGJ5bko3MA&remarks=VFcgQSAt5Y6f55SfSVAgQFREIC0gMjAwTQ&group=WWFoYWhhLUxU
	//ssrs := make([]*SSR, 0)
	//for _, decodedLink := range decodedLinks {
	//    ssr := &SSR{}
	//    re := regexp.MustCompile(`^([^:]+):([^:]+):([^:]+):([^:]+):([^:]+):([^:]+)/?(\S*)$`)
	//}
}
