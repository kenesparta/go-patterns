package builders

type Igloo struct {
	windowType string
	doorType   string
	floor      int
	BuilderInterface
}

func (s *Igloo) WindowType() {
	s.windowType = "Snow Windows"
}

func (s *Igloo) DoorType() {
	s.doorType = "Snow door"
}

func (s *Igloo) NumFloor() {
	s.floor = 1
}

func (s *Igloo) GetHouse() House {
	return House{
		windowType: s.windowType,
		doorType:   s.doorType,
		floor:      s.floor,
	}
}
