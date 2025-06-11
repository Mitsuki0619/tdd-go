package money

type Sum struct {
	augend int
	addend int
}

func (s Sum) Reduce(to string) Money {
	amount := s.augend + s.addend
	return Money{amount: amount, currency: to}
}
