package builders

type BuilderInterface interface {
	WindowType()
	DoorType()
	NumFloor()
	GetHouse() House
	PrintHouse()
}

type HouseType int

const (
	NormalType HouseType = 1 << iota
	IglooType
)

func GetBuilder(h HouseType) BuilderInterface {
	switch h {
	case NormalType:
		return &Normal{}
	case IglooType:
		return &Igloo{}
	default:
		return nil
	}
}
