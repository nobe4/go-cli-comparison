package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nobe4/go-cli-comparison/internal/spec"
)

func main() {
	cmd := cobra.Command{}

	// A
	aFlag := cmd.Flags().BoolP("a", "a", false, "")

	// B
	bFlag := cmd.Flags().CountP("b", "b", "")

	// C
	cFlag := cmd.Flags().StringP("c", "c", "", "")

	// D
	dFlag := cmd.Flags().StringArrayP("d", "d", nil, "")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run the command: %q", err)
		os.Exit(1)

		return
	}

	o := spec.Options{
		A: *aFlag,
		B: *bFlag,
		C: *cFlag,
		D: *dFlag,
	}

	out, err := o.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal options: %q", err)
		os.Exit(1)

		return
	}

	fmt.Fprintln(os.Stdout, string(out))
}
