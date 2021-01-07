package consolehelper

import (
	"io/ioutil"
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
	// TODO: Revise once consolesize is test-safe.
	/*
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

	// TODO: Revise once consolesize is test-safe.
	/*
		cons.UpdateDimensions()

		if cons.Rows == 0 {
			t.Errorf("New instance of ConsoleHelper with Rows == 0")
		}
		if cons.Cols == 0 {
			t.Errorf("New instance of ConsoleHelper with Cols == 0")
		}
	*/
}

func TestFprintf(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var charsWritten int
	var writeErr error
	for _, testStr := range testStrings {
		str = cons.Sprint(testStr)
		strLen = len(str)
		charsWritten, writeErr = cons.Fprintf(ioutil.Discard, testStr)
		if writeErr != nil {
			t.Errorf("Could not write to ioutil.Discard: %s", writeErr)
		}
		if charsWritten != strLen {
			t.Errorf("Wrote %d out of %d chars", charsWritten, strLen)
		}
	}
}

func TestFprint(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var charsWritten int
	var writeErr error
	for _, testStr := range testStrings {
		str = cons.Sprint(testStr)
		strLen = len(str)
		charsWritten, writeErr = cons.Fprint(ioutil.Discard, testStr)
		if writeErr != nil {
			t.Errorf("Could not write to ioutil.Discard: %s", writeErr)
		}
		if charsWritten != strLen {
			t.Errorf("Wrote %d out of %d chars", charsWritten, strLen)
		}
	}
}

func TestFprintln(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var charsWritten int
	var writeErr error
	for _, testStr := range testStrings {
		str = cons.Sprintln(testStr)
		strLen = len(str)
		charsWritten, writeErr = cons.Fprintln(ioutil.Discard, testStr)
		if writeErr != nil {
			t.Errorf("Could not write to ioutil.Discard: %s", writeErr)
		}
		if charsWritten != strLen {
			t.Errorf("Wrote %d out of %d chars", charsWritten, strLen)
		}
	}
}

func TestPrintf(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var charsWritten int
	var writeErr error
	for _, testStr := range testStrings {
		str = cons.Sprint(testStr)
		strLen = len(str)
		charsWritten, writeErr = cons.Printf(testStr)
		if writeErr != nil {
			t.Errorf("Could not write to ioutil.Discard: %s", writeErr)
		}
		if charsWritten != strLen {
			t.Errorf("Wrote %d out of %d chars", charsWritten, strLen)
		}
	}
}

func TestPrint(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var charsWritten int
	var writeErr error
	for _, testStr := range testStrings {
		str = cons.Sprint(testStr)
		strLen = len(str)
		charsWritten, writeErr = cons.Print(testStr)
		if writeErr != nil {
			t.Errorf("Could not write to ioutil.Discard: %s", writeErr)
		}
		if charsWritten != strLen {
			t.Errorf("Wrote %d out of %d chars", charsWritten, strLen)
		}
	}
}

func TestPrintln(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var charsWritten int
	var writeErr error
	for _, testStr := range testStrings {
		str = cons.Sprintln(testStr)
		strLen = len(str)
		charsWritten, writeErr = cons.Println(testStr)
		if writeErr != nil {
			t.Errorf("Could not write to ioutil.Discard: %s", writeErr)
		}
		if charsWritten != strLen {
			t.Errorf("Wrote %d out of %d chars", charsWritten, strLen)
		}
	}
}

func TestSprintf(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var newlinesExpected int
	var newlinesFound int
	for _, testStr := range testStrings {
		strLen = len(testStr)
		if strLen <= cons.Cols {
			newlinesExpected = 0
		} else {
			newlinesExpected = strLen / cons.Cols
		}
		str = cons.Sprintf("%s", testStr)
		newlinesFound = strings.Count(str, newlineSeq)
		if newlinesExpected != newlinesFound {
			t.Errorf("%d character string wrapped on a %d cols console (%d newlines)", strLen, cons.Cols, newlinesFound)
		}
	}
}

func TestSprint(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var newlinesExpected int
	var newlinesFound int
	for _, testStr := range testStrings {
		strLen = len(testStr)
		if strLen <= cons.Cols {
			newlinesExpected = 0
		} else {
			newlinesExpected = strLen / cons.Cols
		}
		str = cons.Sprint(testStr)
		newlinesFound = strings.Count(str, newlineSeq)
		if newlinesExpected != newlinesFound {
			t.Errorf("%d character string wrapped on a %d cols console (%d newlines)", strLen, cons.Cols, newlinesFound)
		}
	}
}

func TestSprintln(t *testing.T) {
	var cons ConsoleHelper
	cons.Rows = testRows
	cons.Cols = testCols

	var str string
	var strLen int
	var newlinesExpected int
	var newlinesFound int
	for _, testStr := range testStrings {
		strLen = len(testStr)
		if strLen <= cons.Cols {
			newlinesExpected = 1
		} else {
			newlinesExpected = strLen/cons.Cols + 1
		}
		str = cons.Sprintln(testStr)
		newlinesFound = strings.Count(str, newlineSeq)
		if newlinesExpected != newlinesFound {
			t.Errorf("%d character string wrapped on a %d cols console (expected %d, found %d newlines)\n%s", strLen, cons.Cols, newlinesExpected, newlinesFound, str)
		}
	}
}
