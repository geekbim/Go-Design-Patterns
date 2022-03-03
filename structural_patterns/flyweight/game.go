package main

type game struct {
	terrorist        []*player
	counterTerrorist []*player
}

func newGame() *game {
	return &game{
		terrorist:        make([]*player, 1),
		counterTerrorist: make([]*player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorist = append(g.terrorist, player)
	return
}

func (g *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorist = append(g.counterTerrorist, player)
	return
}
