package testHelper

import (
	"reflect"
	"testing"
)

func AssertEqualsPrimitive[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Assertion failed: Expected %v, but got %v", expected, actual)
	}
}

func AssertEquals(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assertion failed (DeepEqual): Expected %v, but got %v", expected, actual)
	}
}

func AssertTrue(t *testing.T, b bool) {
	t.Helper()
	if !b {
		t.Errorf("it is false")
	}
}

func AssertFalse(t *testing.T, b bool) {
	t.Helper()
	if b {
		t.Errorf("it is true")
	}
}
