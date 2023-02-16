package main

func main(){
	mac := &mac{}
	windows := &windows{}
	windowsAdapter := &windowsAdapter{window: windows}
	client := &client{}
	client.insertSquareUsbIntoComputer(mac)
	client.insertSquareUsbIntoComputer(windowsAdapter)
}