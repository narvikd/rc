package fmtutils

import (
	"fmt"
	"github.com/TwiN/go-color"
	"os"
)

// Fatalln wraps FatallnMsg, it prints err.Error().
func Fatalln(err error) {
	FatallnMsg(err.Error())
}

// FatallnMsg prints a message in red and then exits.
func FatallnMsg(msg string) {
	fmt.Println(color.InRed(msg))
	os.Exit(1)
}
