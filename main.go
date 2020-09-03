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

	parser := map[string]func(string, bool) (Shadow, error){
		"ssd": parseSSD,
		"ss":  parseSS,
		"ssr": parseSSR,
	}[*originType]

	if parser == nil {
		log.Fatalln("wrong type")
	}

	rawShadow, err := parser(*origin, *cache)
	if err != nil {
		log.Fatalln("error:", err)
	}
	list := rawShadow.(SSList)
	if *verbose {
		for index, item := range list {
			log.Println(index, item)
		}
	}
	log.Println("total:", len(list))

	config, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(ssToConfig(list))
	if err != nil {
		log.Fatalln("marshal error:", err)
	}
	ioutil.WriteFile(*originType+"_config.json", []byte(config), 0755)
}
