package abstract_factory

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) Wheels() int {
	return 2
}

func (c *CruiseMotorbike) Seats() int {
	return 2
}

func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
