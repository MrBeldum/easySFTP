package driver

import "testing"

func TestInformativeRepeatsRaisesTwoToThree(t *testing.T) {
	// Issue #227: two is the broken release-matrix default; leave 1 and 3+ alone.
	if got := informativeRepeats(2); got != 3 {
		t.Errorf("informativeRepeats(2) = %d, want 3", got)
	}
	for _, n := range []int{1, 3, 5} {
		if got := informativeRepeats(n); got != n {
			t.Errorf("informativeRepeats(%d) = %d, want unchanged", n, got)
		}
	}
}
