package main

import (
	"flag"
	"fmt"
)

func main() {
	var s string

	flag.StringVar(&s, "name", "", "use well")

	flag.Parse()
	
	fmt.Println(s)
}
