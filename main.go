package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	originType := flag.String("type", "ss", `type of the origin, can be "ss" or "ssr"`)
	origin := flag.String("origin", "", `origin url`)
	cache := flag.Bool("cache", true, `whether to load data from cache file`)
	verbose := flag.Bool("verbose", false, `whether to show detail`)
	flag.Parse()

	shadowsocksList, err := parseSS(*origin, *cache)
	if err != nil {
		log.Fatalln("error:", err)
	}
	if *verbose {
		for index, item := range shadowsocksList {
			log.Println(index, item)
		}
	}
	log.Println("total:", len(shadowsocksList))

	config, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(ssToConfig(shadowsocksList))
	if err != nil {
		log.Fatalln("marshal error:", err)
	}
	ioutil.WriteFile(*originType+"_config.json", []byte(config), 0755)
}
