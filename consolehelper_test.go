package consolehelper

import (
	"strings"
	"testing"
)

const (
	testRows     int    = 15
	testCols     int    = 15
	testString08 string = "8 chars."
	testString14 string = "14 characters."
	testString92 string = "A sentence with 92 characters that should split over multiple lines thanks to the soft wrap."
	newlineSeq   string = "\n"
)

var (
	testStrings [3]string
)

func init() {
	testStrings[0] = testString08
	testStrings[1] = testString14
	testStrings[2] = testString92
}

func TestConsoleHelper(t *testing.T) {
	var cons ConsoleHelper

	if cons.AutoUpdate != false {
		t.Errorf("New ConsoleHelper variable with AutoUpdate != 0")
	}
	if cons.Rows != 0 {
		t.Errorf("New ConsoleHelper variable with Rows != 0")
	}
	if cons.Cols != 0 {
		t.Errorf("New ConsoleHelper variable with Cols != 0")
	}
}

func TestNew(t *testing.T) {
	cons, _ := New()

	if cons.AutoUpdate != false {
		t.Errorf("New instance of ConsoleHelper with AutoUpdate != false")
	}
	/*
		// TODO: Revise once consolesize is test-safe.
		if cons.Rows == 0 {
			t.Errorf("New instance of ConsoleHelper with Rows == 0")
		}
		if cons.Cols == 0 {
			t.Errorf("New instance of ConsoleHelper with Cols == 0")
		}
	*/
}

func TestUpdateDimensions(t *testing.T) {
	var cons ConsoleHelper

	if cons.Rows != 0 {
		t.Errorf("New ConsoleHelper variable with Rows != 0")
	}
	if cons.Cols != 0 {
		t.Errorf("New ConsoleHelper variable with Cols != 0")
	}

	/*
		// TODO: Revise once consolesize is test-safe.
		cons.UpdateDimensions()

		if cons.Rows == 0 {
			t.Errorf("New instance of ConsoleHelper with Rows == 0")
		}
		if cons.Cols == 0 {
			t.Errorf("New instance of ConsoleHelper with Cols == 0")
		}
	*/
}

func TestSprintf(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var newlinesExpected bool
	var newlines int
	var newlinesFound bool
	for _, testStr := range testStrings {
		strLen = len(testStr)
		if strLen <= cons.Cols {
			newlinesExpected = false
		} else {
			newlinesExpected = true
		}
		str = cons.Sprintf("%s", testStr)
		newlines = strings.Count(str, newlineSeq)
		if newlines == 0 {
			newlinesFound = false
		} else {
			newlinesFound = true
		}
		if newlinesExpected != newlinesFound {
			t.Errorf("%d character string wrapped on a %d cols console (%d newlines)", strLen, cons.Cols, newlines)
		}
	}
}

func TestSprint(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var newlinesExpected bool
	var newlines int
	var newlinesFound bool
	for _, testStr := range testStrings {
		strLen = len(testStr)
		if strLen <= cons.Cols {
			newlinesExpected = false
		} else {
			newlinesExpected = true
		}
		str = cons.Sprint(testStr)
		newlines = strings.Count(str, newlineSeq)
		if newlines == 0 {
			newlinesFound = false
		} else {
			newlinesFound = true
		}
		if newlinesExpected != newlinesFound {
			t.Errorf("%d character string wrapped on a %d cols console (%d newlines)", strLen, cons.Cols, newlines)
		}
	}
}
