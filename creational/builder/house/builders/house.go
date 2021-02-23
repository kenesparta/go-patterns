package builders

import "fmt"

type House struct {
	windowType string
	doorType   string
	floor      int
}

func (h *House) PrintHouse() {
	fmt.Printf("\nHouse Door Type: %s\n", h.doorType)
	fmt.Printf("House Window Type: %s\n", h.windowType)
	fmt.Printf("House Num Floor: %d\n", h.floor)
}