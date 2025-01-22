package spec

import (
	"errors"
	"strconv"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    Options
		o2   Options
		want bool
	}{
		// This ensures that default values are equal to each others, so later
		// tests can focus on comparing fields independently.
		{
			want: true,
		},

		// A
		{
			o:    Options{A: true},
			o2:   Options{A: false},
			want: false,
		},

		// B
		{
			o:    Options{B: 0},
			o2:   Options{B: 1},
			want: false,
		},

		// C
		{
			o:    Options{C: "a"},
			o2:   Options{C: "b"},
			want: false,
		},

		// D
		{
			o:    Options{D: []string{"a"}},
			o2:   Options{D: []string{"a", "b"}},
			want: false,
		},
		{
			o:    Options{D: []string{"b", "a"}},
			o2:   Options{D: []string{"a", "b"}},
			want: false,
		},
		{
			o:    Options{D: []string{"a", "a", "b"}},
			o2:   Options{D: []string{"a", "b"}},
			want: false,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			got := test.o.Equal(test.o2)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func errIsNil(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s        string
		o        Options
		want     bool
		checkErr func(t *testing.T, err error)
	}{
		// Ensures that empty strings are failing correctly.
		{
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !errors.Is(err, errInvalidOptions) {
					t.Errorf("got %v, want %v", err, errInvalidOptions)
				}
			},
		},

		// This ensures that an empty json string unmarshals correctly to an
		// empty Option, so later tests can focus on comparing fields
		// independently.
		{
			s:    "{}",
			want: true,
		},

		// A
		{
			s:    `{ "a": true }`,
			o:    Options{A: true},
			want: true,
		},

		// B
		{
			s:    `{ "b": 2 }`,
			o:    Options{B: 2},
			want: true,
		},

		// C
		{
			s:    `{ "c": "a" }`,
			o:    Options{C: "a"},
			want: true,
		},

		// D
		{
			s:    `{ "d": ["a", "b"] }`,
			o:    Options{D: []string{"a", "b"}},
			want: true,
		},
		{
			s:    `{ "d": ["b", "a"] }`,
			o:    Options{D: []string{"a", "b"}},
			want: false,
		},

		// Additional tests

		// Order of the field does not matter.
		{
			s: `{
                "d": ["a", "b"],
                "a": true
            }`,
			o: Options{
				A: true,
				D: []string{"a", "b"},
			},
			want: true,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if test.checkErr == nil {
				test.checkErr = errIsNil
			}

			o, err := Unmarshal(test.s)
			test.checkErr(t, err)

			if err == nil && o.Equal(test.o) != test.want {
				t.Errorf("got %v, want %v", o, test.o)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o        Options
		want     bool
		checkErr func(t *testing.T, err error)
	}{
		// This ensures that an empty json string unmarshals correctly to an
		// empty Option, so later tests can focus on comparing fields
		// independently.
		{
			o:    Options{},
			want: true,
		},

		// A
		{
			o:    Options{A: true},
			want: true,
		},

		// B
		{
			o:    Options{B: 0},
			want: true,
		},

		// C
		{
			o:    Options{C: "a"},
			want: true,
		},

		// D
		{
			o:    Options{D: []string{"a", "b"}},
			want: true,
		},

		// Additional tests
		// With many fields
		{
			o: Options{
				A: true,
				D: []string{"a", "b"},
			},
			want: true,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			if test.checkErr == nil {
				test.checkErr = errIsNil
			}

			s, err := test.o.Marshal()
			test.checkErr(t, err)

			o2, err := Unmarshal(s)
			if err != nil {
				t.Errorf("could not unmarshal the marshalled Options: %v", err)
			}

			if test.o.Equal(o2) != test.want {
				t.Errorf("got %v, want %v", test.o, o2)
			}
		})
	}
}
