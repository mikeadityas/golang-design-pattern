package abstract_factory

type LuxuryCar struct{}

func (l *LuxuryCar) Wheels() int {
	return 4
}

func (l *LuxuryCar) Seats() int {
	return 5
}

func (l *LuxuryCar) Doors() int {
	return 4
}
