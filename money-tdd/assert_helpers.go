package money

import (
	"reflect"
	"testing"
)

func assertEqualsPrimitive[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("Assertion failed: Expected %v, but got %v", expected, actual)
	}
}

func assertEquals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assertion failed (DeepEqual): Expected %v, but got %v", expected, actual)
	}
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Errorf("it is false")
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Errorf("it is true")
	}
}