package utils

import (
	"log"
	url1 "net/url"

	"github.com/golang/protobuf/jsonpb"
)

func decodeTarget(target string) (cache bool, u string, opt *options, err error) {
	cache, u, opt, err = false, "", &options{}, nil

	url, err := url1.Parse(target)
	if err != nil {
		return
	}
	query := url.Query()
	if query.Get("cache") == "true" {
		cache = true
	}
	if len(query.Get("loglevel")) > 0 {
		opt.loglevel = query.Get("loglevel")
	}
	query.Del("cache")
	query.Del("loglevel")
	url.RawQuery = query.Encode()
	u = url.String()
	return
}

func GetSubscriptionSS(target string) ([]byte, error) {
	cache, u, opt, err := decodeTarget(target)
	if err != nil {
		return nil, err
	}
	log.Printf("Normalized: %v\n", u)

	shadowsocksList, err := parseSS(u, cache)
	log.Printf("Loaded %d ss configs\n", len(shadowsocksList))
	if err != nil {
		return nil, err
	}
	config, err := (&jsonpb.Marshaler{}).MarshalToString(ssToConfig(shadowsocksList, opt))
	if err != nil {
		return nil, err
	}
	return []byte(config), nil
}

func GetShadows(target string) (shadows []string, err error) {
	shadows, err = []string{}, nil

	_, u, _, err := decodeTarget(target)
	if err != nil {
		return
	}
	shadowsocksList, err := parseSS(u, true)
	if err != nil {
		return
	}

	for _, shadowsocks := range shadowsocksList {
		shadows = append(shadows, shadowsocks.remarks)
	}
	return
}
