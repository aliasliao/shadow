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
	verbose := flag.Bool("verbose", false, `whether to show detail`)
	flag.Parse()

	if *originType == "ss" {
		res, err := getSSD(*origin, *cache)
		if err != nil {
			log.Fatalln("error:", err)
		}
		if *verbose {
		    for index, item := range res {
		    	log.Println(index, *item)
			}
		}
		log.Println("total:", len(res))
	} else if *originType == "ssr" {
		res, err := getSSR(*origin, *cache)
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
