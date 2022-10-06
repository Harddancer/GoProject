package paser

import (
	"fmt"
	"os"
	"parseravito/applications"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func WFile() string {
	text, _ := applications.Client()
	file, err := os.Create("avto.html")

	if err != nil {
		fmt.Println("Файл не создан:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Файл создан.")
	return text
}

func Grab() <-chan string { //функция вернет канал, из которого мы будем читать данные типа string
	c := make(chan string)
	go func() {
		for { //в вечном цикле собираем данные
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(WFile()))
			if err == nil {
				if s := strings.TrimSpace(doc.Find(".iva-item-titleStep-pdebR").Text()); s != "" {
					c <- s //и отправляем их в канал
				}
			}
			time.Sleep(100 * time.Millisecond)
		}

	}()

	return c
}
