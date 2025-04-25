package main

import "fmt"

func main() {
	addidasFac, _ := GetFactory("Addidas")
	nikeFac, _ := GetFactory("Nike")

	addiadasShoe := addidasFac.MakeShoe()
	addidasShirt := addidasFac.MakeShirt()

	nikeShoe := nikeFac.MakeShoe()
	nikeShirt := nikeFac.MakeShirt()

	printShoeDetails(addiadasShoe)
	printShirtDetails(addidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(shoe ShoeProduct) {
	fmt.Printf("Shoe Logo: %s, Size: %d\n", shoe.GetLogo(), shoe.GetSize())
}

func printShirtDetails(shirt ShirtProduct) {
	fmt.Printf("Shirt Logo: %s, Size: %d\n", shirt.GetLogo(), shirt.GetSize())
}
