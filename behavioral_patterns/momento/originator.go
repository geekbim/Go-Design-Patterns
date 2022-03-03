package main

type originator struct {
	state string
}

func (o *originator) createMomento() *momento {
	return &momento{state: o.state}
}

func (o *originator) restoreMomento(m *momento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}
