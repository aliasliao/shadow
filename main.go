package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/aliasliao/shadow/utils"
)

func main() {
	originType := flag.String("type", "ss", `type of the origin, can be "ss" or "ssr"`)
	origin := flag.String("origin", "", `origin url`)
	flag.Parse()

	config, err := utils.GetSubscriptionSS(*origin)
	if err != nil {
		log.Fatalln("marshal error:", err)
	}
	ioutil.WriteFile("build/"+*originType+"_config.json", config, 0755)
}
