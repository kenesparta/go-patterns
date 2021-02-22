package main

type Adidas struct {
	SportsFactoryInterface
}

func (a *Adidas) MakeShoe() ShoeInterface {
	shoe := new(Shoe)
	*shoe.fLogo() = "Adidas Shoe"
	*shoe.fSize() = 50
	return &AdidasShoe{*shoe}
}

func (a *Adidas) MakeShirt() ShirtInterface {
	shirt := new(Shirt)
	*shirt.fLogo() = "Adidas Shirt"
	*shirt.fSize() = 50
	return &AdidasShirt{*shirt}
}
