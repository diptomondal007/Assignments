package bridge_practice

type Computer interface {
	Print()
	SetPrinter(p Printer)
}
