package main

import "fmt"

type SportProductsFactory interface {
	MakeShoe() ShoeProduct
	MakeShirt() ShirtProduct
}

func GetFactory(brand string) (SportProductsFactory, error) {
	switch brand {
	case "Nike":
		return &NikeFactory{}, nil
	case "Addidas":
		return &AddidasFactory{}, nil
	default:
		return nil, fmt.Errorf("Unknown brand: %s", brand)
	}
}
