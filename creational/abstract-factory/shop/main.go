package main

func main() {
	adidasFactory, _ := GetSportsFactory(AdidasType)
	nikeFactory, _ := GetSportsFactory(NikeType)

	// Adidas
	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	// Nike
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe.fPrintDetails()
	adidasShirt.fPrintDetails()

	nikeShoe.fPrintDetails()
	nikeShirt.fPrintDetails()
}
