package Builder

type IBuildProcess interface {
	SetWheels() IBuildProcess
	SetSeats() IBuildProcess
	SetStructure() IBuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	builder IBuildProcess
}

func (m *ManufacturingDirector) Construct(){
	//...implementation goes here
	m.builder.SetSeats().SetStructure().SetWheels()
}

func (m *ManufacturingDirector) SetBuilder(b IBuildProcess){
	//...implementation goes here
	m.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

func (c *CarBuilder) SetWheels() IBuildProcess{
	c.v.Wheels = 4
	return  c
}

func (c *CarBuilder) SetSeats() IBuildProcess{
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() IBuildProcess{
	c.v.Structure = "Car"
	return c
}


type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

func (b *BikeBuilder) SetWheels() IBuildProcess{
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() IBuildProcess{
	b.v.Seats =2
	return b
}

func (b *BikeBuilder) SetStructure() IBuildProcess{
	b.v.Structure = "Motorbike"
	return b
}