package main

import (
	"flag"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"log"
)

func main() {
	originType := flag.String("type", "ss", `type of the origin, can be "ss" or "ssr"`)
	origin := flag.String("origin", "", `origin url`)
	cache := flag.Bool("cache", true, `whether to load data from cache file`)
	verbose := flag.Bool("verbose", false, `whether to show detail`)
	flag.Parse()

	if *originType == "ssd" {
		res, err := parseSSD(*origin, *cache)
		if err != nil {
			log.Fatalln("error:", err)
		}
		if *verbose {
			for index, item := range res {
				log.Println(index, *item)
			}
		}
		log.Println("total:", len(res))

		config, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(ssToConfig(res))
		if err != nil {
			log.Fatalln("marshal error:", err)
		}
		ioutil.WriteFile("ssd_config.json", []byte(config), 0755)
	} else if *originType == "ss" {
		res, err := parseSS(*origin, *cache)
		if err != nil {
			log.Fatalln("error:", err)
		}
		if *verbose {
			for index, item := range res {
				log.Println(index, *item)
			}
		}
		log.Println("total:", len(res))

		config, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(ssToConfig(res))
		if err != nil {
			log.Fatalln("marshal error:", err)
		}
		ioutil.WriteFile("ss_config.json", []byte(config), 0755)
	} else if *originType == "ssr" {
		res, err := parseSSR(*origin, *cache)
		if err != nil {
			log.Fatalln("error:", err)
		}
		if *verbose {
			for index, item := range res {
				log.Println(index, *item)
			}
		}
		log.Println("total:", len(res))
	} else {
		log.Fatalln("wrong type")
	}
}
