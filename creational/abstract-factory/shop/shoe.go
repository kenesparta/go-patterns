package main

import "fmt"

type Shoe struct {
	logo string
	size int
	ShoeInterface
}

func (s *Shoe) fLogo() *string {
	return &s.logo
}

func (s *Shoe) fSize() *int {
	return &s.size
}

func (s *Shoe) fPrintDetails() {
	fmt.Printf("Shoe Logo: %s", s.logo)
	fmt.Println()
	fmt.Printf("Shoe Size: %d", s.size)
	fmt.Println()
}
