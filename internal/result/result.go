package result

import (
	"bytes"
)

type Result [][]bool

func New(h, w int) Result {
	results := make([][]bool, h)

	for i := range h {
		results[i] = make([]bool, w)
	}

	return results
}

func (r Result) Equal(o Result) bool {
	if len(r) != len(o) {
		return false
	}

	if len(r[0]) != len(o[0]) {
		return false
	}

	for i, row := range r {
		for j, cell := range row {
			if cell != o[i][j] {
				return false
			}
		}
	}

	return true
}

func (r Result) Marshal() []byte {
	lines := [][]byte{}

	for _, row := range r {
		line := []byte{}

		for _, cell := range row {
			if cell {
				line = append(line, '1')
			} else {
				line = append(line, '0')
			}
		}

		lines = append(lines, line)
	}

	return bytes.Join(lines, []byte{'\n'})
}

func Unmarshal(in []byte, r *Result) {
	lines := bytes.Split(in, []byte{'\n'})

	o := New(len(lines), len(lines[0]))

	for i, line := range lines {
		for j, cell := range line {
			o[i][j] = cell == '1'
		}
	}

	*r = o
}
