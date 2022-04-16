package main

import (
	"fmt"
	"os"
)

func handleArgs() {
	if len(os.Args) < 4 {
		initFail()
	}
	if os.Args[1] != "g" || os.Args[2] != "c" || os.Args[3] == "" {
		initFail()
	}
}

func initFail() {
	fmt.Println("Usage of rc:")
	fmt.Println(" rc g c ComponentName")
	fmt.Println("	Create a React Component. (Required)")
	os.Exit(1)
}
