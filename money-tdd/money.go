package money

type Money struct {
	amount   int
	currency string
}

func NewDollar(amount int) Money {
	return Money{amount: amount, currency: "USD"}
}

func NewFranc(amount int) Money {
	return Money{amount: amount, currency: "CHF"}
}

func (d Money) Amount() int {
	return d.amount
}

func (d Money) Currency() string {
	return d.currency
}

func (d Money) Reduce() Money {
	return d
}

func (d Money) Plus(addend Money) Expression {
	return Sum{augend: d.amount, addend: addend.amount}
}

func (d Money) Times(m int) Money {
	return Money{d.Amount() * m, d.Currency()}
}

func (d Money) Equals(d2 Money) bool {
	if d.amount == d2.amount && d.Currency() == d2.Currency() {
		return true
	}
	return false
}
