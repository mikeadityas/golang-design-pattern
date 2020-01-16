package builder

type ManufacturingDirector interface {
	SetBuilder(BuildProcess)
	Construct()
}

var manufacturingDirectorInstance *manufacturingDirector

func GetManufacturingDirectorInstance() ManufacturingDirector {
	if manufacturingDirectorInstance == nil {
		manufacturingDirectorInstance = new(manufacturingDirector)
	}
	return manufacturingDirectorInstance
}

type manufacturingDirector struct {
	builder BuildProcess
}

func (d *manufacturingDirector) SetBuilder(b BuildProcess) {
	d.builder = b
}

func (d *manufacturingDirector) Construct() {
	d.builder.SetWheels().SetSeats().SetStructure()
}
