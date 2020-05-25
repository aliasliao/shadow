package main

import (
	"flag"
	"log"
)

// shadow -type ss -origin https://example.com -cache
func main() {
	originType := flag.String("type", "ss", `type of the origin, can be "ss" or "ssr"`)
	origin := flag.String("origin", "", `origin url`)
	cache := flag.Bool("cache", false, `whether to load data from cache file`)
	flag.Parse()
	if *originType == "ss" {
		res, err := getSSD(*origin, *cache)
		if err != nil {
			log.Fatalln("error:", err)
		}
		log.Println("total:", len(res))
	} else if *originType == "ssr" {
		res, err := getSSR(*origin, *cache)
		if err != nil {
			log.Fatalln("error:", err)
		}
		log.Println("total:", len(res))
	} else {
		log.Fatalln("wrong type")
	}
}
