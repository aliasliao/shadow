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

	parser := map[string]func(string, bool) ([]Shadow, error){
		"ssd": parseSSD,
		"ss":  parseSS,
		"ssr": parseSSR,
	}[*originType]

	if parser == nil {
		log.Fatalln("wrong type")
	}

	shadows, err := parser(*origin, *cache)
	if err != nil {
		log.Fatalln("error:", err)
	}
	if *verbose {
		for index, shadow := range shadows {
			log.Println(index, shadow)
		}
	}
	log.Println("total:", len(shadows))
	var values []*Shadowsocks
	for _, shadow := range shadows {
		value := shadow.(*Shadowsocks) // TODO: assert by type
		values = append(values, value)
	}

	config, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(ssToConfig(values))
	if err != nil {
		log.Fatalln("marshal error:", err)
	}
	ioutil.WriteFile(*originType+"_config.json", []byte(config), 0755)
}
