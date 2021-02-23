package director

import "house/builders"

type Director struct {
	builder builders.BuilderInterface
}

func NewDirector(b builders.BuilderInterface) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b builders.BuilderInterface) {
	d.builder = b
}

func (d *Director) BuildHouse() builders.House {
	d.builder.DoorType()
	d.builder.WindowType()
	d.builder.NumFloor()
	return d.builder.GetHouse()
}
