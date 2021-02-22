package main

type Nike struct {
	SportsFactoryInterface
}

func (a *Nike) MakeShoe() ShoeInterface {
	shoe := new(Shoe)
	*shoe.fLogo() = "Nike Shoe"
	*shoe.fSize() = 74
	return &NikeShoe{*shoe}
}

func (a *Nike) MakeShirt() ShirtInterface {
	shirt := new(Shirt)
	*shirt.fLogo() = "Nike Shirt"
	*shirt.fSize() = 15
	return &NikeShirt{*shirt}
}
