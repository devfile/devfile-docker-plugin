package tests

import (
	"reflect"
	"testing"
)

func ExpectEqual(t *testing.T, actual interface{}, expected interface{}, message string) bool {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s. \nexpected = %v \n     got = %v", message, expected, actual)
		return false
	}
	return true
}

func ExpectNil(t *testing.T, value interface{}, message string) bool {
	t.Helper()
	if value != nil {
		t.Errorf("%s. expected '%v' to be nil", message, value)
		return false
	}
	return true
}

func ExpectNotNil(t *testing.T, value interface{}, message string) bool {
	t.Helper()
	if value == nil {
		t.Errorf("%s. expected '%v' to not be nil", message, value)
		return false
	}
	return true
}

func ToBoolPtr(b bool) *bool {
	return &b
}
