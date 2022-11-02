package main

import (
	"fmt"
	"parsernioss/workers"
)

func main() {
	m := workers.Readernioss("mrcnioss.csv")
	mg, mb := workers.Createdata("sbcmrc.txt", m)
	fmt.Println(workers.Writegoodvoice(mg))
	fmt.Println(workers.Writebadvoice(mb))

}
