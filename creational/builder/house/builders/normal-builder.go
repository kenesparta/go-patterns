package builders

import "fmt"

type Normal struct {
	windowType string
	doorType   string
	floor      int
	BuilderInterface
}

func (s *Normal) WindowType() {
	s.windowType = "Wooden Window"
}

func (s *Normal) DoorType() {
	s.doorType = "Wooden Door"
}

func (s *Normal) NumFloor() {
	s.floor = 5
}

func  (s *Normal) PrintHouse()  {
	fmt.Printf("Normal House Door Type: %s\n", s.doorType)
	fmt.Printf("Normal House Window Type: %s\n", s.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", s.floor)
}

func (s *Normal) GetHouse() House {
	return House{
		windowType: s.windowType,
		doorType:   s.doorType,
		floor:      s.floor,
	}
}
