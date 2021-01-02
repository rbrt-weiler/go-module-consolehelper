package consolehelper

import "testing"

func TestConsoleHelper(t *testing.T) {
	var cons ConsoleHelper

	if cons.AutoUpdate != false {
		t.Errorf("New instance of ConsoleHelper with AutoUpdate != 0")
	}
	if cons.Rows != 0 {
		t.Errorf("New instance of ConsoleHelper with Rows != 0")
	}
	if cons.Cols != 0 {
		t.Errorf("New instance of ConsoleHelper with Cols != 0")
	}
}
