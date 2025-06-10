package money

import "testing"

func TestMultiplication(t *testing.T) {
	dollar := NewDollar(2)
	assertEquals(t, dollar.Times(5), NewDollar(10))
	assertEquals(t, dollar.Times(3), NewDollar(6))
}

func TestEquality(t *testing.T) {
	assertTrue(t, NewDollar(3).Equals(NewDollar(3)))
	assertFalse(t, NewDollar(4).Equals(NewDollar(5)))
	assertTrue(t, NewFranc(3).Equals(NewFranc(3)))
	assertFalse(t, NewFranc(4).Equals(NewFranc(5)))
	assertFalse(t, NewDollar(3).Equals(NewFranc(3)))
}

func TestCurrency(t *testing.T) {
	assertEqualsPrimitive(t, NewDollar(1).Currency(), "USD")
	assertEqualsPrimitive(t, NewFranc(1).Currency(), "CHF")
}
