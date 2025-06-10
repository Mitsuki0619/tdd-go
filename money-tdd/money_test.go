package money

import "testing"

func TestMultiplication(t *testing.T) {
	dollar := NewDollar(2)
	assertEqual(t, dollar.Times(5), NewDollar(10))
	assertEqual(t, dollar.Times(3), NewDollar(6))
}

func TestEquality(t *testing.T) {
	assertTrue(t, NewDollar(3).Equals(NewDollar(3)))
	assertFalse(t, NewDollar(4).Equals(NewDollar(5)))
}

func TestFrancMultiplication(t *testing.T) {
	franc := NewFranc(2)
	assertEqual(t, franc.Times(5), NewFranc(10))
	assertEqual(t, franc.Times(3), NewFranc(6))
}
