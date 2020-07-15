package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

//-------------------------------------
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newmsg string) {
	newmsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newmsg)
	return
}

//---------------------------------------------
type NewPrinter interface {
	PrintStored() string
}

type PrintAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrintAdapter) PrintStored() (newmsg string) {
	if p.OldPrinter != nil {
		newmsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newmsg = p.OldPrinter.Print(newmsg)
	} else {
		newmsg = p.Msg
	}
	return
}
