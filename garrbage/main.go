package main

import "fmt"

func main() {
	s := []string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}

	var galact Meta
	galact.Printing(s)

}

type Meta struct {
	number int
	planet string
}

func (m *Meta) Printing(s []string) {
	for i, v := range s {
		m.number = i
		m.planet = v
		fmt.Printf("%d %s\n", m.number, m.planet)
	}
}
