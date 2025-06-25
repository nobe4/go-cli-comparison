package spec

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var errInvalidOptions = errors.New("invalid options")

type State int

const (
	StateSuccess State = iota
	StatePartial
	StateFailure
)

type Spec struct {
	ID          int
	Description string
	Default     string
	Type        string
	Regex       string
	Comments    []string
	State       State

	Tests []Test
}
type Specs []Spec

type TestResults []bool

type Results []TestResults

func (s Specs) Populate(r io.Reader) error {

	c, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	results := Results{}

	if err := json.Unmarshal(c, &results); err != nil {
		return err
	}

	fmt.Printf("%+v\n", results)

	return nil
}
