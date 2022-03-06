package main

type offCommand struct {
	device device
}

func (o *offCommand) execute() {
	o.device.off()
}
