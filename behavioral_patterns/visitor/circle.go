package main

type circle struct {
	radius int
}

func (s *circle) accept(v visitor) {
	v.visitForCircle(s)
}

func (s *circle) getType() string {
	return "Circle"
}
