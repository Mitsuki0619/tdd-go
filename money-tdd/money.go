package money

type Money interface {
	Times(m int) Money
	Equals(d Money) bool
	Amount() int
}

type Dollar struct {
	amount int
}

func NewDollar (amount int) Dollar {
	return Dollar{amount: amount}
}

func (d Dollar) Times (m int) Money  {
	return NewDollar(d.Amount() * m)
}

func (d Dollar) Equals (d2 Money) bool {
	if d.Amount() == d2.Amount() {
		return true
	}
	return false
}

func (d Dollar) Amount () int {
	return d.amount
}