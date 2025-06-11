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
	assertFalse(t, NewDollar(3).Equals(NewFranc(3)))
}

func TestCurrency(t *testing.T) {
	assertEqualsPrimitive(t, NewDollar(1).Currency(), "USD")
	assertEqualsPrimitive(t, NewFranc(1).Currency(), "CHF")
}

func TestSimpleAddition(t *testing.T) {
	dollar := NewDollar(3)
	sum := dollar.Plus(NewDollar(4))
	bank := Bank{}
	reduced := bank.Reduce(sum, "USD")
	assertEquals(t, NewDollar(7), reduced)
}

func TestReduceSum(t *testing.T) {
	sum := Sum{augend: 4, addend: 6}
	bank := Bank{}
	result := bank.Reduce(sum, "USD")
	assertEquals(t, result, NewDollar(10))
}
