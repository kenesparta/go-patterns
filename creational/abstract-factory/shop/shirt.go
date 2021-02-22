package main

import "fmt"

type Shirt struct {
	logo string
	size int
	ShirtInterface
}

func (s *Shirt) fLogo() *string {
	return &s.logo
}

func (s *Shirt) fSize() *int {
	return &s.size
}

func (s *Shirt) fPrintDetails() {
	fmt.Printf("Shirt Logo: %s", s.logo)
	fmt.Println()
	fmt.Printf("Shirt Size: %d", s.size)
	fmt.Println()
}
