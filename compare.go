// Package testcmp compares values in tests using [github.com/google/go-cmp/cmp].
package testcmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Compare logs a [cmp.Diff] when [cmp.Equal] reports that got and want differ.
// Options are passed to both Equal and Diff.
func Compare(t *testing.T, got, want any, opts ...cmp.Option) {
	t.Helper()
	if !cmp.Equal(got, want, opts...) {
		t.Log(cmp.Diff(got, want, opts...))
	}
}
