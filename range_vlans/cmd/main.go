package main

import (
	"fmt"
	"vlans/vlans"
)

func main() {

	gv := vlans.GroupVlans([]int{1, 2, 3, 18, 19, 20, 21, 22, 23, 400, 437, 567, 588})
	fmt.Println(gv)
}
