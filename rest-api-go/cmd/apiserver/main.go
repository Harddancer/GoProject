package main

import (
	"log"

	apiserver "github.com/Harddancer/GoProject/rest-api-go/internal/app/apiserver"
)

func main() {

	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)

	}

}
