package main

import (
	"log"

	michan "http-rest-api/apiserver"
)

func main() {

	s := michan.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
