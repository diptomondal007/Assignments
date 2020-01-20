package abstractfactory

type SportBike struct {
	
}

func (s SportBike) NumWheels() int {
	return 2
}

func (s SportBike) NumSeats() int {
	return 1
}

func (s *SportBike) GetMotorBikeType() int{
	return SportMotorBikeType
}

