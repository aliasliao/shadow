package main

import "fmt"

func main() {
	ySSR := "http://sub.yahaha.pro/link/TF1WoRrrADeyTYTu?mu=0"
	ssrs, err := getSSR(ySSR, true)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("total:", len(ssrs))
}
