package assert

import (
	"errors"
	"fmt"
	"testing"
)

// EqErr asserts that errors exp and is have the same content.
func EqErr(t *testing.T, exp error, is error, description string) {
	t.Helper()

	if exp == nil {
		Nil(t, is, description)

		return
	}

	if is == nil {
		t.Errorf(errorMsgFormat, failed, description, is)

		return
	}

	if fmt.Sprintf("%v", exp) != fmt.Sprintf("%v", is) {
		t.Errorf(errorMsgFormat, failed, description, is)
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}

// IsErr asserts that errors exp is contained in given error message.
func IsErr(t *testing.T, exp error, given error, description string) {
	t.Helper()

	if exp == nil {
		Nil(t, given, description)

		return
	}

	if given == nil {
		t.Errorf(errorMsgFormat, failed, description, given)

		return
	}

	// given: wrapped error chain.
	// exp: static error.
	if errors.Is(given, exp) {
		// expected error is included in given error chain.
		t.Logf(successMsgFormat, succeed, description)
	} else {
		// expected error is not included in given error chain.
		t.Errorf(errorMsgFormat, failed, description, given)
	}
}
