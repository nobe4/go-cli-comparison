package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nobe4/go-cli-comparison/internal/spec"
)

func main() {
	aFlag := flag.Bool("a", false, "")

	flag.Parse()

	o := spec.Options{
		A: *aFlag,
	}

	out, err := o.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal options: %q", err)
		os.Exit(1)

		return
	}

	fmt.Fprintln(os.Stdout, string(out))
}
