package utils

import (
	"log"
	url "net/url"

	"github.com/golang/protobuf/jsonpb"
)

func GetSubscriptionSS(target string) ([]byte, error) {
	cache := false
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if query.Get("cache") == "true" {
		cache = true
	}
	query.Del("cache")
	u.RawQuery = query.Encode()
	log.Printf("Normalized: %v\n", u.String())

	shadowsocksList, err := parseSS(u.String(), cache)
	log.Printf("Loaded %d ss configs\n", len(shadowsocksList))
	if err != nil {
		return nil, err
	}
	config, err := (&jsonpb.Marshaler{}).MarshalToString(ssToConfig(shadowsocksList))
	if err != nil {
		return nil, err
	}
	return []byte(config), nil
}
