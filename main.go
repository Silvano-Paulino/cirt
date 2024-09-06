package main

import (
	"fmt"
	"os"

	"github.com/silvano-paulino/cmd"
)

func main() {

	cirt := cmd.NewCirt()
	if err := cirt.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
