package abstractfactory

//FamilyCar ...
type FamilyCar struct{}

//NumWheels ...
func (f *FamilyCar) NumWheels() int {
	return 4
}

//NumSeats ...
func (f *FamilyCar) NumSeats() int {
	return 5
}

//NumDoors ...
func (f *FamilyCar) NumDoors() int {
	return 5
}
