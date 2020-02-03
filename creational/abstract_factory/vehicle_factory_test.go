package abstract_factory

import (
	"fmt"
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal("A factory of type 'Motorbike' must exist")
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal("A motorbike factory must be able to generate 'Sport Motorbike'")
	}

	sportWheels := motorbikeVehicle.Wheels()
	if sportWheels != 2 {
		t.Error("A sport motorbike must have 2 wheels")
	}

	sportSeats := motorbikeVehicle.Seats()
	if sportSeats != 1 {
		t.Error("A sport motorbike must only have 1 seat")
	}

	sportMotorbike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Sport motorbike must be a motorbike")
	}

	sportMotorbikeType := sportMotorbike.GetMotorbikeType()
	if sportMotorbikeType != SportMotorbikeType {
		t.Error(fmt.Sprintf("Sport motorbike must have type of %d", SportMotorbikeType))
	}

	motorbikeVehicle2, err := motorbikeF.NewVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal("A motorbike factory must be able to generate 'Cruise Motorbike'")
	}

	cruiseWheels := motorbikeVehicle2.Wheels()
	if cruiseWheels != 2 {
		t.Error("A cruise motorbike must have 2 wheels")
	}

	cruiseSeats := motorbikeVehicle2.Seats()
	if cruiseSeats != 2 {
		t.Error("A cruise motorbike must have 2 seats")
	}

	cruiseMotorbike, ok := motorbikeVehicle2.(Motorbike)
	if !ok {
		t.Fatal("Cruise motorbike must be a motorbike")
	}

	cruiseMotorbikeType := cruiseMotorbike.GetMotorbikeType()
	if cruiseMotorbikeType != CruiseMotorbikeType {
		t.Error(fmt.Sprintf("Cruise motorbike must have type of %d", CruiseMotorbikeType))
	}
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal("A factory of type 'Car' must exist")
	}

	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal("A car factory must be able to generate 'Luxury Car'")
	}

	luxuryWheels := carVehicle.Wheels()
	if luxuryWheels != 4 {
		t.Error("A luxury car must have 4 wheels")
	}

	luxurySeats := carVehicle.Seats()
	if luxurySeats != 5 {
		t.Error("A luxury car must have 5 seats")
	}

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Luxury car must be a car")
	}

	luxuryCarDoors := luxuryCar.Doors()
	if luxuryCarDoors != 4 {
		t.Error("A luxury car must have 4 doors")
	}

	carVehicle2, err := carF.NewVehicle(FamilyCarType)
	if err != nil {
		t.Fatal("A car factory must be able to generate 'Family Car'")
	}

	familyWheels := carVehicle2.Wheels()
	if familyWheels != 4 {
		t.Error("A family car must have 4 wheels")
	}

	familySeats := carVehicle2.Seats()
	if familySeats != 7 {
		t.Error("A family car must have 7 seats")
	}

	familyCar, ok := carVehicle2.(Car)
	if !ok {
		t.Fatal("Family car must be a car")
	}

	familyCarDoors := familyCar.Doors()
	if familyCarDoors != 4 {
		t.Error("A family car must have 4 doors")
	}
}
