package spec

import (
	"encoding/json"
	"fmt"
)

// See ./list.go for details on how those are used.
type Options struct {
	A bool     `json:"a"`
	B int      `json:"b"`
	C string   `json:"c"`
	D []string `json:"d"`
}

func (o Options) Equal(o2 Options) bool {
	if o.A != o2.A {
		return false
	}

	if o.B != o2.B {
		return false
	}

	if o.C != o2.C {
		return false
	}

	if len(o.D) != len(o2.D) {
		return false
	}

	for i, v := range o.D {
		if v != o2.D[i] {
			return false
		}
	}

	return true
}

func Unmarshal(s []byte) (Options, error) {
	o := Options{}

	if err := json.Unmarshal(s, &o); err != nil {
		return o, fmt.Errorf("failed to unmarshal %w: %w", errInvalidOptions, err)
	}

	return o, nil
}

func (o Options) Marshal() ([]byte, error) {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return b, fmt.Errorf("failed to marshal %w: %w", errInvalidOptions, err)
	}

	return b, nil
}

func (o Options) String() string {
	b, err := o.Marshal()
	if err != nil {
		return fmt.Sprintf("failed to marshal: %v", err)
	}

	return string(b)
}
