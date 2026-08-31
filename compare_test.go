package testcmp_test

import (
	"testing"

	"github.com/charlieparkes/go-testcmp"
	"github.com/google/go-cmp/cmp"
)

func TestCompareEqual(t *testing.T) {
	t.Parallel()
	testcmp.Compare(t, 1, 1)
}

func TestComparePassesOptionsToEqualAndDiff(t *testing.T) {
	t.Parallel()
	calls := 0
	opt := cmp.Comparer(func(_, _ int) bool {
		calls++
		return false
	})
	testcmp.Compare(t, 1, 1, opt)
	if calls < 2 {
		t.Fatalf("Comparer called %d times, want at least 2 (Equal and Diff)", calls)
	}
}
