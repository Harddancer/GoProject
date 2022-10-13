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
				var sp string
				doc.Find(".iva-item-titleStep-pdebR>a").Each(func(i int, s *goquery.Selection) {
					//fmt.Println(s.Attr("title"))
					p, _ := s.Attr("title")
					sp = fmt.Sprintf("%s///%s", sp, p)

				})
				c <- sp //запись в канал

			}
			time.Sleep(10 * time.Millisecond)
		}

	}()

	return c
}
