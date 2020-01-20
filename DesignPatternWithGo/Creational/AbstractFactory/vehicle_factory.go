package abstractfactory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}


const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

const (
	SportMotorBikeType = 1
	CruiseMotorBikeType = 2
)

const (
	CarFactoryType = 1
	MotorBikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error){
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorBikeFactoryType:
		return new(MotorBikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n",f))
	}
}

type CarFactory struct{}
type MotorBikeFactory struct {}

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not " +
			"recognized\n", v))
	}
}

func (m *MotorBikeFactory) Build(v int) (Vehicle, error){
	switch v {
	case SportMotorBikeType:
		return new(SportBike), nil
	case CruiseMotorBikeType:
		return new(CruiseBike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
