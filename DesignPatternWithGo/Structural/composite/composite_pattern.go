package main

import "fmt"

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type Athlete struct{}

func (a *Athlete) Train() { fmt.Println("Training")
}

type SwimmerImpl struct {

}

func(s *SwimmerImpl) Swim(){
	fmt.Println("Swim")
}

type SwimmerCompositeB struct {
	Trainer
	Swimmer
}

func main(){
	swimmer := SwimmerCompositeB{
		Trainer: &Athlete{},
		Swimmer: &SwimmerImpl{},
	}

	swimmer.Swim()
	swimmer.Train()
}