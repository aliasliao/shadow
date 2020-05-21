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
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(raw)))
	_, err := base64.StdEncoding.Decode(bytes, raw)
	if err != nil {
		fmt.Println("warning", err)
	}
	return bytes
}

func getSSR(url string, cache bool) (string, error) {
	var raw []byte
	CacheFile := os.TempDir() + "/" + base64.RawStdEncoding.EncodeToString([]byte(url))
	if _, err := os.Stat(CacheFile); cache && !os.IsNotExist(err) {
		fmt.Println("Read from cache...")
		raw, _ = ioutil.ReadFile(CacheFile)
	} else {
		res, err := (&http.Client{Timeout: 10 * time.Second}).Get(url)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		raw, _ = ioutil.ReadAll(res.Body)
		ioutil.WriteFile(CacheFile, raw, 0755)
	}
	fmt.Printf("len(raw) = %d, cache file = %s\n", len(raw), CacheFile)
	links := strings.Split(string(safeDecode(raw)), "\n")
	decodedLinks := make([]string, 0)
	for _, link := range links {
		decodedLinks = append(decodedLinks, string(safeDecode([]byte(link[6:]))))
	}
	fmt.Println(decodedLinks)

	//// ssr://120.232.150.53:65533:auth_aes128_md5:chacha20-ietf:tls1.2_ticket_auth:cXdlcnQ/?obfsparam=ZTIwMDgyNzUzOS5kb3dubG9hZC53aW5kb3dzdXBkYXRlLmNvbQ&protoparam=Mjc1Mzk6RGJ5bko3MA&remarks=VFcgQSAt5Y6f55SfSVAgQFREIC0gMjAwTQ&group=WWFoYWhhLUxU
	//ssrs := make([]*SSR, 0)
	//for _, decodedLink := range decodedLinks {
	//    ssr := &SSR{}
	//    re := regexp.MustCompile(`^([^:]+):([^:]+):([^:]+):([^:]+):([^:]+):([^:]+)/?(\S*)$`)
	//}

	return string(raw), nil
}
