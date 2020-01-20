package adapter_practice


type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (w *WindowsAdapter) InsertInSquarePort(){
	w.WindowsMachine.InsertInCirclePort()
}