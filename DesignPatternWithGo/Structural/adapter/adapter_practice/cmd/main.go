package main

import "designpatternwithgo/Structural/adapter/adapter_practice"

func main(){
	client := &adapter_practice.Client{}
	mac := &adapter_practice.Mac{}
	client.InsertSquareUsbIntoComputer(mac)
	windowsMachine := &adapter_practice.Windows{}
	windowsMachineAdapter := &adapter_practice.WindowsAdapter{
		WindowsMachine:windowsMachine,
	}
	client.InsertSquareUsbIntoComputer(windowsMachineAdapter)

}
