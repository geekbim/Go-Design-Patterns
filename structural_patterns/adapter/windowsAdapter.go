package main

import "fmt"

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
