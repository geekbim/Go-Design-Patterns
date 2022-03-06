package main

type onCommand struct {
	device device
}

func (o *onCommand) execute() {
	o.device.on()
}
