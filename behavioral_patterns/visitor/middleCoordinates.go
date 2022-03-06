package main

import "fmt"

type middleCoordinates struct {
	x int
	y int
}

func (m *middleCoordinates) visitForSquare(s *square) {
	// Calculate area for square
	// Then assign in to the area instance variable
	fmt.Println("Calculating middle point coordinates for square")
}

func (m *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (m *middleCoordinates) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
