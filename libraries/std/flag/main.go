package main

import (
	"flag"
	"fmt"
)

func main() {
	aFlag := flag.Int("a", 0, "flag a")

	flag.Parse()

	//nolint: forbidigo // This is fine for now
	fmt.Println("aFlag:", *aFlag)
}
