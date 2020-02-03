package abstract_factory

type SportMotorbike struct{}

func (s *SportMotorbike) Wheels() int {
	return 2
}

func (s *SportMotorbike) Seats() int {
	return 1
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
