package abstractfactory

import (
	"testing"
)

func TestMotorBikeFactory(t *testing.T){
	motorbikeF, err := BuildFactory(MotorBikeFactoryType)
	if err != nil{
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.Build(SportMotorBikeType)
	if err != nil{
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())
	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok{
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorBikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err) }
	carVehicle, err := carF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err) }
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed") }
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors()) }