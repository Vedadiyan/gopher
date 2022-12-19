package main

import (
	"os"

	flaggy "github.com/vedadiyan/flaggy/pkg"
)

func main() {
	flags := Flags{}
	flaggy.Parse(&flags, os.Args[1:])
}
