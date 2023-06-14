package apiserver

import "fmt"

// Api server

type Apiserver struct{}

// New
func New() *Apiserver {
	fmt.Println("Done")
	return &Apiserver{}
}

// Start ...
func (s Apiserver) Start() error {
	return nil

}
