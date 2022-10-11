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

func Grab() <-chan string { //string
	c := make(chan string)
	go func() {
		for {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(WFile()))
			if err == nil {
				if s := doc.Find(".iva-item-titleStep-pdebR > a").Text(); s != "" {
					fmt.Printf("%#v\n\n", s)
					c <- s //запись в канал
				}
			}
			time.Sleep(10 * time.Millisecond)
		}

	}()

	return c
}
