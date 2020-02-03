package abstract_factory

type FamilyCar struct{}

func (f *FamilyCar) Wheels() int {
	return 4
}

func (f *FamilyCar) Seats() int {
	return 7
}

func (f *FamilyCar) Doors() int {
	return 4
}
