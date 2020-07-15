package bridge

import (
	"errors"
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use API implementation: Message %s\n", err)
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Contest received on Writer was empty")
	return
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}
	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct. \n Expected: %s\n Actual:%s\n", expectedErrorMessage, err)
		}
	}
	testwriter := TestWriter{}
	api2 = PrinterImpl2{
		Writer: &testwriter,
	}
	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err)
	}
	if testwriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer.\n Expected: %s\nActual: %s\n", expectedMessage, testwriter.Msg)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"
	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}
	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testwriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testwriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if testwriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesnt match.\n Expected: %s\nActual: %s\n", expectedMessage, testwriter.Msg)
	}

}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"
	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}
	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testwriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testwriter,
		},
	}
	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if testwriter.Msg != expectedMessage {
		t.Errorf("The expected message on io.Writer doesnt match\n Expect: %s\n Actual: %s\n", expectedMessage, testwriter.Msg)
	}
}
