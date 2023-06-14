package apiserver

// Api server

type Apiserver struct{}

// New
func New() *Apiserver {
	return &Apiserver{}
}

// Start ...
func (s Apiserver) Start() error {
	return nil

}
