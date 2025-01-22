package result

import (
	"bytes"
	"testing"
)

func TestMarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		r    Result
		want []byte
	}{
		{
			r:    Result{[]bool{false}},
			want: []byte("0"),
		},
		{
			r:    Result{[]bool{false, true}, []bool{true, false}},
			want: []byte("01\n10"),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := test.r.Marshal()

			if !bytes.Equal(got, test.want) {
				t.Fatalf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		i    []byte
		want Result
	}{
		{
			i:    []byte("0"),
			want: Result{[]bool{false}},
		},
		{
			i:    []byte("01\n10"),
			want: Result{[]bool{false, true}, []bool{true, false}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := Result{}
			Unmarshal(test.i, &got)

			if !got.Equal(test.want) {
				t.Fatalf("want %v, got %v", test.want, got)
			}
		})
	}
}
