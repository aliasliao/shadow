package main

import "fmt"

func main() {
	ySSR := "http://sub.yahaha.pro/link/TF1WoRrrADeyTYTu?mu=0"
	_, err := getSSR(ySSR, true)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
