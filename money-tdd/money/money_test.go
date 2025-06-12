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
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	testHelper.AssertEquals(t, NewDollar(7), reduced)
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(4), NewDollar(6))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	testHelper.AssertEquals(t, result, NewDollar(10))
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	testHelper.AssertEquals(t, NewDollar(1), result)
}

func TestIdentityRate(t *testing.T) {
	testHelper.AssertEquals(t, 1, NewBank().Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	testHelper.AssertEquals(t, NewDollar(10), result)
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)
	result := bank.Reduce(sum, "USD")
	testHelper.AssertEquals(t, NewDollar(15), result)
}

func TestSumTimes(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveBucks, tenFrancs).Times(2)
	result := bank.Reduce(sum, "USD")
	testHelper.AssertEquals(t, NewDollar(20), result)
}
