package main

type igloBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIgloBuilder() *igloBuilder {
	return &igloBuilder{}
}

func (b *igloBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *igloBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *igloBuilder) setNumFloor() {
	b.floor = 1
}

func (b *igloBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
