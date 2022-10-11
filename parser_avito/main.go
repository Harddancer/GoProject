package main

import (
	"fmt"
	"parseravito/paser"
	"regexp"
	"strings"
)

func main() {

	channel := paser.Grab()
	s := <-channel
	var auto Avto
	ch := strings.Split(s, ",")
	auto.Printing(ch)

}

type Avto struct {
	number int
	vendor string
	model  string
	year   string
}

func (a *Avto) Printing(ch []string) {
	for i, v := range ch {

		re := regexp.MustCompile("[0-9]+")
		y := strings.Join(re.FindAllString(v, 1), "")
		a.number = i
		a.year = y

		re2 := regexp.MustCompile("[a-zA-Z]+\\s")
		vend := strings.Join(re2.FindAllString(v, 1), "")
		a.vendor = vend

		v = strings.TrimPrefix(v, " ")

		re3 := regexp.MustCompile("\\s+[a-x,A-Z,а-я,0-9,-,()]+")
		mod := strings.Join(re3.FindAllString(v, 1), "")
		a.model = mod

		fmt.Println(a.number, a.year, a.vendor, a.model)
	}

}
