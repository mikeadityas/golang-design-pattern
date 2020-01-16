package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := GetManufacturingDirectorInstance()

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and was %d\n", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure of a car must be 'Car' and was %s\n", car.Structure)
	}

	manufacturingComplex2 := GetManufacturingDirectorInstance()

	if manufacturingComplex2 != manufacturingComplex {
		t.Error("ManufacturingDirector is expected to be a Singleton")
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex2.SetBuilder(bikeBuilder)
	manufacturingComplex2.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Seats != 1 {
		t.Errorf("Seats on a bike must be 1 and was %d\n", bike.Seats)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure of a bike must be 'Bike' and was %s\n", bike.Structure)
	}
}
