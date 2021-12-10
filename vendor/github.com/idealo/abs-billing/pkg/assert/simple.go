package assert

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

const (
	succeed           = "\u2713"
	failed            = "\u2717"
	successMsgFormat  = "\t%s: %s."
	errorMsgFormat    = "\t%s: %s: %v"
	errorEqMsgFormat  = "\t%s: %s: expected: '%v' got: '%v' (differences: [%v])"
	floatPrecision100 = 100
)

// True asserts that 'condition' is true.
func True(t *testing.T, condition bool, description string) {
	t.Helper()

	if !condition {
		t.Errorf(errorMsgFormat, failed, description, condition)
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}

// False asserts that 'condition' is false.
func False(t *testing.T, condition bool, description string) {
	t.Helper()

	if condition {
		t.Errorf(errorMsgFormat, failed, description, condition)
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}

// Nil asserts that 'is' is nil.
func Nil(t *testing.T, is interface{}, description string) {
	t.Helper()

	if !isNil(is) {
		t.Errorf(errorMsgFormat, failed, description, is)
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}

// NNil asserts that 'is' isnt nil.
func NNil(t *testing.T, is interface{}, description string) {
	t.Helper()

	if isNil(is) {
		t.Errorf(errorMsgFormat, failed, description, is)
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}

// Eq asserts that 'exp' matches 'is' using reflection.
func Eq(t *testing.T, exp interface{}, is interface{}, description string) {
	t.Helper()

	// without increasing the precision for float variables the following comparison will return nil
	// deep.Equal([]float64{1, 15.9}, []float64{1, 15.8999999999998})
	deep.FloatPrecision = floatPrecision100
	if result := deep.Equal(exp, is); result != nil {
		t.Errorf(errorEqMsgFormat, failed, description, exp, is, strings.Join(result, ","))
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}

// NEq asserts that 'exp' doesn't match 'is' using reflection.
func NEq(t *testing.T, exp interface{}, is interface{}, description string) {
	t.Helper()

	// without increasing the precision for float variables the following comparison will return nil
	// deep.Equal([]float64{1, 15.9}, []float64{1, 15.8999999999998})
	deep.FloatPrecision = floatPrecision100
	if result := deep.Equal(exp, is); result == nil {
		t.Errorf(errorMsgFormat, failed, description, is)
	} else {
		t.Logf(successMsgFormat, succeed, description)
	}
}
