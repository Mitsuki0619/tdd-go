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

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return Money{m.amount / rate, to}
}

func (m Money) Plus(addend Money) Expression {
	return Sum{augend: m.amount, addend: addend.amount}
}

func (m Money) Times(multiplier int) Money {
	return Money{m.Amount() * multiplier, m.Currency()}
}

func (m Money) Equals(d2 Money) bool {
	if m.amount == d2.amount && m.Currency() == d2.Currency() {
		return true
	}
	return false
}
