package main

import "fmt"

const (
	AdidasType BrandType = 1 << iota
	NikeType
	UmbroType
)

const errorMessage = "NOT IMPLEMENTED"

// BrandType The type of the brand on the Sports Factory
type BrandType int

// SportsFactoryInterface Abstract factory interface
type SportsFactoryInterface interface {
	MakeShoe() ShoeInterface
	MakeShirt() ShirtInterface
}

func GetSportsFactory(brand BrandType) (SportsFactoryInterface, error) {
	switch brand {
	case AdidasType:
		return &Adidas{}, nil
	case NikeType:
		return &Nike{}, fmt.Errorf(errorMessage)
	case UmbroType:
		return nil, fmt.Errorf(errorMessage)
	default:
		return nil, fmt.Errorf(errorMessage)
	}
}
