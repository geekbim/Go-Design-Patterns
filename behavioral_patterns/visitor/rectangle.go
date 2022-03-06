package main

type rectangle struct {
	l int
	b int
}

func (s *rectangle) accept(v visitor) {
	v.visitForRectangle(s)
}

func (s *rectangle) getType() string {
	return "Rectangle"
}
