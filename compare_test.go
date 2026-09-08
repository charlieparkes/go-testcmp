package testcmp_test

import (
	"testing"

	"github.com/charlieparkes/go-testcmp"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type point struct {
	X, Y int
}

func TestCompareEqual(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		got  any
		want any
		opts []cmp.Option
	}{
		"ints":        {got: 1, want: 1},
		"strings":     {got: "hello", want: "hello"},
		"bools":       {got: true, want: true},
		"int slices":  {got: []int{1, 2, 3}, want: []int{1, 2, 3}},
		"string maps": {got: map[string]int{"a": 1, "b": 2}, want: map[string]int{"a": 1, "b": 2}},
		"structs":     {got: point{X: 1, Y: 2}, want: point{X: 1, Y: 2}},
		"approx floats with option": {
			got:  1.0,
			want: 1.0000001,
			opts: []cmp.Option{cmpopts.EquateApprox(0, 1e-6)},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testcmp.Compare(t, tc.got, tc.want, tc.opts...)
		})
	}
}
