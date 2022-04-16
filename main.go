package main

import (
	"log"
	"os"
	"rc/generate/component"
)

func main() {
	handleArgs()
	createComponent()
}

func createComponent() {
	for i := 3; i < len(os.Args); i++ {
		err := component.Create(os.Args[i])
		if err != nil {
			log.Fatalln(err)
		}
	}
}
