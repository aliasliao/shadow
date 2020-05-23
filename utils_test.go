package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSafeDecode(t *testing.T) {
	t.Run("decode 4s", func(t *testing.T) {
		input := []byte("R2hNVlhoVlVr")
		expect := []byte("GhMVXhVUk")
		if output := safeDecode(input); bytes.Compare(output, expect) != 0 {
			t.Errorf("Expect: %s, got: %s", expect, output)
		}
	})
	t.Run("decode none 4s", func(t *testing.T) {
		input := []byte("R2hNVlhoVlVrRQo")
		expect := []byte("GhMVXhVUkE\n")
		if output := safeDecode(input); bytes.Compare(output, expect) != 0 {
			t.Errorf("Expect: %s, got: %s", expect, output)
		}
	})
}

func TestDecodeLink(t *testing.T) {
	input := "ssr://dGtwYS5oay1mLnlhaGFoYS5wcm86NjU1MzM6YXV0aF9hZXMxMjhfbWQ1OmNoYWNoYTIwLWlldGY6dGxzMS4yX3RpY2tldF9hdXRoOmVXRm9ZV2hoYkhSay8_b2Jmc3BhcmFtPVpUSXdNRGd5TnpVek9TNWtiM2R1Ykc5aFpDNTNhVzVrYjNkemRYQmtZWFJsTG1OdmJRJnByb3RvcGFyYW09TWpjMU16azZSR0o1YmtvM01BJnJlbWFya3M9U0VzZ1JpQXRJREZIWW5CeiZncm91cD1XV0ZvWVdoaExVeFVSQQ"
	expect := &SSR{
		server:        "tkpa.hk-f.yahaha.pro",
		serverPort:    65533,
		localAddress:  "",
		localPort:     0,
		timeout:       0,
		workers:       0,
		password:      safeDecodeStr("eWFoYWhhbHRk"),
		method:        "chacha20-ietf",
		obfs:          "tls1.2_ticket_auth",
		obfsParam:     safeDecodeStr("ZTIwMDgyNzUzOS5kb3dubG9hZC53aW5kb3dzdXBkYXRlLmNvbQ"),
		protocol:      "auth_aes128_md5",
		protocolParam: safeDecodeStr("Mjc1Mzk6RGJ5bko3MA"),
	}
	output, err := decodeLink(input)
	if err != nil {
		t.Error(err)
	}
	serializedExpect, serializedOutput := fmt.Sprintf("%v", expect), fmt.Sprintf("%v", output)
	if serializedExpect != serializedOutput {
		t.Errorf("\nExpect: %s\nOutput: %s\n", serializedExpect, serializedOutput)
	}
}
