package main

import (
	"house/builders"
	"house/director"
)

func main() {
	normalBuilder := builders.GetBuilder(builders.NormalType)
	iglooBuilder := builders.GetBuilder(builders.IglooType)

	dir := director.NewDirector(normalBuilder)
	normalHouse := dir.BuildHouse()
	normalHouse.PrintHouse()

	dir.SetBuilder(iglooBuilder)
	iglooHouse := dir.BuildHouse()
	iglooHouse.PrintHouse()
}
