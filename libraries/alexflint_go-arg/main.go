/*
Name = alexflint/go-arg
URL  = https://github.com/alexflint/go-arg
*/
package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/nobe4/go-cli-comparison/internal/spec"
)

func main() {
	var args struct {
		A bool
		// B unimplemented
		C string
		D []string `arg:"-d,separate"`
	}

	arg.MustParse(&args)

	o := spec.Options{
		A: args.A,
		C: args.C,
		D: args.D,
	}

	out, err := o.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal options: %q", err)
		os.Exit(1)

		return
	}

	fmt.Fprintln(os.Stdout, string(out))
}
