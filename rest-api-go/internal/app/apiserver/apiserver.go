package apiserver

import "fmt"

// Api server

type Apiserver struct {
	config *Config
}

// New
func New(config *Config) *Apiserver {
	fmt.Println("Done")
	return &Apiserver{
		config: config,
	}
}

// Start ...
func (s Apiserver) Start() error {
	return nil

}
