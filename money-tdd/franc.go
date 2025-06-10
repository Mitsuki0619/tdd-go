package money

type Franc struct {
	amount int
}

func NewFranc (amount int) Franc {
	return Franc{amount: amount}
}

func (d Franc) Times (m int) Money  {
	return NewFranc(d.Amount() * m)
}

func (d Franc) Equals (d2 Money) bool {
	if d.Amount() == d2.Amount() {
		return true
	}
	return false
}

func (d Franc) Amount () int {
	return d.amount
}