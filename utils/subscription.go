package utils

import "github.com/golang/protobuf/jsonpb"

func GetSubscriptionSS(url string) ([]byte, error) {
	shadowsocksList, err := parseSS(url, false)
	if err != nil {
		return nil, err
	}
	config, err := (&jsonpb.Marshaler{}).MarshalToString(ssToConfig(shadowsocksList))
	if err != nil {
		return nil, err
	}
	return []byte(config), nil
}
