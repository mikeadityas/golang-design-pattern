package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("Factory with ID %d not recognized\n", f),
		)
	}
}
