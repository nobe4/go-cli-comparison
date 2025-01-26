/*
Name = flag
URL  = https://pkg.go.dev/flag
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nobe4/go-cli-comparison/internal/spec"
)

func main() {
	// A
	aFlag := flag.Bool("a", false, "")

	// B
	bFlag := 0

	flag.BoolFunc("b", "", func(string) error {
		bFlag++

		return nil
	})

	// C
	cFlag := flag.String("c", "", "")

	// D
	dFlag := []string{}

	flag.Func("d", "", func(s string) error {
		dFlag = append(dFlag, s)

		return nil
	})

	flag.Parse()

	o := spec.Options{
		A: *aFlag,
		B: bFlag,
		C: *cFlag,
		D: dFlag,
	}

	out, err := o.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal options: %q", err)
		os.Exit(1)

		return
	}

	fmt.Fprintln(os.Stdout, string(out))
}
