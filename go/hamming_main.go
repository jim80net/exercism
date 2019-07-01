package main

import (
	"fmt"
	"os"

	"./hamming"
)

func main() {
	distance, err := hamming.Distance(os.Args[1], os.Args[2])
	fmt.Printf("distance: %#v, error: %v", distance, err)
}
