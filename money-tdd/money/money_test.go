package money

import (
	testHelper "money-tdd/test_helper"
	"testing"
)

func TestMultiplication(t *testing.T) {
	dollar := NewDollar(2)
	testHelper.AssertEquals(t, dollar.Times(5), NewDollar(10))
	testHelper.AssertEquals(t, dollar.Times(3), NewDollar(6))
}

func TestEquality(t *testing.T) {
	testHelper.AssertTrue(t, NewDollar(3).Equals(NewDollar(3)))
	testHelper.AssertFalse(t, NewDollar(4).Equals(NewDollar(5)))
	testHelper.AssertFalse(t, NewDollar(3).Equals(NewFranc(3)))
}

func TestCurrency(t *testing.T) {
	testHelper.AssertEqualsPrimitive(t, NewDollar(1).Currency(), "USD")
	testHelper.AssertEqualsPrimitive(t, NewFranc(1).Currency(), "CHF")
}

func TestSimpleAddition(t *testing.T) {
	dollar := NewDollar(3)
	sum := dollar.Plus(NewDollar(4))
	bank := Bank{}
	reduced := bank.Reduce(sum, "USD")
	testHelper.AssertEquals(t, NewDollar(7), reduced)
}

func TestReduceSum(t *testing.T) {
	sum := Sum{augend: 4, addend: 6}
	bank := Bank{}
	result := bank.Reduce(sum, "USD")
	testHelper.AssertEquals(t, result, NewDollar(10))
}
