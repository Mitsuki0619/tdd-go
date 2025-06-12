package money

type Pair struct {
	from string
	to   string
}

func NewPair(from, to string) Pair {
	return Pair{from, to}
}

func (p Pair) Equals(p2 Pair) bool {
	return p.from == p2.from && p.to == p2.to
}

func (p Pair) HashCode() int {
	return 0
}
