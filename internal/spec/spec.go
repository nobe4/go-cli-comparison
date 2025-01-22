package spec

import (
	"encoding/json"
	"errors"
	"fmt"
)

var errInvalidOptions = errors.New("invalid options")

type Options struct {
	// desc:      Boolean toggle, default to false.
	// default:   0
	// type:      bool
	// regex:     -\w
	// e.g:       -a
	A bool `json:"a"`

	// desc:      Count how many times the flag was passed, default to 0;
	// default:   0
	// type:      int
	// regex:     (-\w)( \1)*
	// e.g.:      -b -b -b
	B int `json:"b"`

	// desc:      Store an string
	// default:   ""
	// type:      string
	// regex:     -\w(=| )?\.+
	// e.g.:      -c=abcd
	C string `json:"c"`

	// desc:      Store an array of strings
	// default:   []
	// type:      []string
	// regex:     (-\w)(=| )?\.+( \1(=| )?\.+)*
	// e.g.:      -d=a -d b -dc
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
