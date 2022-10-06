package main

import (
	"fmt"
	"parseravito/paser"
	"strings"
)

func main() {
	channel := paser.Grab()
	// type Avto struct {
	// 	number int
	// 	vendor string
	// 	// model  string
	// 	// year   string
	// }
	// var auto Avto
	// var ch []string
	s := <-channel
	// ch = strings.Split(s, " ")
	fmt.Print(strings.TrimSpace(s))
	//fmt.Println(ch)
	// for _, v := range ch {
	// 	v2 := strings.Split(v, " ")
	// 	// auto.number = i
	// 	// auto.vendor = v2[0]
	// 	// auto.model = v2[1]
	// 	// auto.year = v2[2]
	// 	fmt.Println(v2)

	// }

}
