package main

import (
	"fmt"
	"parseravito/paser"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	channel := paser.Grab()
	s := <-channel

	var auto Avto
	ch := strings.Split(s, "///")
	auto.Printing(ch)

}

type Avto struct {
	number int
	vendor string
	model  string
	year   string
	price  int
}

func (a *Avto) Printing(ch []string) {
	for i, v := range ch {
		// fmt.Println(v)
		re := regexp.MustCompile("[0-9]{4}")
		y := strings.Join(re.FindAllString(v, 1), " ")
		a.number = i
		a.year = y

		re2 := regexp.MustCompile("^\\S+")
		vend := strings.Join(re2.FindAllString(v, 1), "")
		a.vendor = vend

		v = strings.TrimPrefix(v, " ")
		// fmt.Print(v)

		re3 := regexp.MustCompile("\\s\\w+\\,[\\s]")
		mod := strings.Join(re3.FindAllString(v, 1), "")
		a.model = mod

		re4 := regexp.MustCompile("а\\s(.*?\\.)")
		pr := strings.Join(re4.FindAllString(v, 1), "")
		re4_1 := regexp.MustCompile("(\\d)")
		pr2 := re4_1.FindAllString(pr, 10)
		pr3 := strings.Join(pr2, " ")
		pr4 := strings.Replace(pr3, " ", "", -1)
		str_number, err := strconv.Atoi(pr4)
		if err != nil {
			fmt.Print(err)
		} else {
			a.price = str_number
			fmt.Printf("Номер %d Год выпуска %s Вендор %s Модель %s Цена %d в руб.\n", a.number, a.year, a.vendor, a.model, a.price)

		}

	}
}
