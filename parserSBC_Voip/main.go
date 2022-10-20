package main

import (
	"fmt"
	"parsernioss/workers"
)

func main() {
	m := workers.Readernioss("nioss.csv")
	mb, mg := workers.Createdata("sbc.csv", m)
	fmt.Println(workers.Writegoodvoice(mg))
	fmt.Println(workers.Writebadvoice(mb))
}
