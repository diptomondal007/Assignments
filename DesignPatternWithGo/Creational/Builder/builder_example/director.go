package builder_example

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director{
	return &Director{
		builder:b,
	}
}

func (d *Director) SetBuilder(b IBuilder){
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}