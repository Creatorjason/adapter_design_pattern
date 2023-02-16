package main


type windowsAdapter struct{
	window *windows
}
func (w *windowsAdapter) insertIntoSquarePort(){
	w.window.insertIntoCircularPort()
}

