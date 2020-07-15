package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrintAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didnt match: %s\n", returnedMsg)
	}
	adapter = PrintAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("Message didnt match: %s\n", returnedMsg)
	}
}
