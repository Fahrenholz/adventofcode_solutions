package assert

import "testing"

// FailFast makes the test exit if it's failed.
func FailFast(t *testing.T) {
	t.Helper()

	if t.Failed() {
		t.FailNow()
	}
}
