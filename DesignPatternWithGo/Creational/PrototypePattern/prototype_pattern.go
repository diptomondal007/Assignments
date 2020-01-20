package PrototypePattern

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtCloner() ShirtCloner {
	return &ShirtCache{}
}

type ShirtCache struct {}

func(sc *ShirtCache) GetClone(s int) (ItemInfoGetter, error){
	switch s {
	case White:
		newItem := *WhitePrototype
		return &newItem, nil
	case Black:
		newItem := *BlackPrototype
		return &newItem, nil
	case Blue:
		newItem := *BluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs%f\n", s.SKU, s.Color, s.Price)
}

var WhitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "Empty",
	Color: White,
}

var BlackPrototype *Shirt = &Shirt{ Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var BluePrototype *Shirt = &Shirt{ Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (s *Shirt) GetPrice() float32{
	return s.Price
}