package abstractfactory

type LuxuryCar struct {}

func (l *LuxuryCar) NumWheels() int {
	return 4
}

func (l *LuxuryCar) NumSeats() int {
	return 5
}

func (l *LuxuryCar) NumDoors() int{
	return 4
}


