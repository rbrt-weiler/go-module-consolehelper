/*

Package consolehelper aids with pretty printing on the CLI. It provides fmt-like functionality with integrated text wrapping. The border for text wrapping is determined automatically, based on the console size.

	package main

	import (
		"fmt"
		"os"

		consolehelper "gitlab.com/rbrt-weiler/go-module-consolehelper"
	)

	func main() {
		fox := "The quick brown fox jumps over the lazy dog."
		cons, consErr := consolehelper.New()
		if consErr != nil {
			fmt.Printf("Error creating consolehelper instance: %s\n", consErr)
			os.Exit(1)
		}
		cons.Println(fox) // will be wrapped if necessary
		cons.Cols = 15    // fake a 15 column console
		cons.Println(fox) // will be wrapped
		os.Exit(0)
	}

*/
package consolehelper

import (
	"fmt"
	"io"

	text "github.com/jedib0t/go-pretty/v6/text"
	consolesize "github.com/nathan-fiscaletti/consolesize-go"
)

const (
	// ModuleName contains the short name of the module.
	ModuleName string = "consolehelper"
	// ModuleVersion contains the version number of the module.
	ModuleVersion string = "1.0.2"
	// ModuleURL contains the URL where the module can be retrieved.
	ModuleURL string = "https://gitlab.com/rbrt-weiler/go-module-consolehelper"
)

// ConsoleHelper encapsulates functionality for pretty printing on the console.
type ConsoleHelper struct {
	// Set AutoUpdate to true to always refresh the console size with each relevant function call.
	AutoUpdate bool
	// The number of rows provided by the console.
	Rows int
	// The number of columns provided by the console.
	Cols int
}

// New creates a new ConsoleHelper instance.
func New() (c ConsoleHelper, err error) {
	c = ConsoleHelper{}
	err = nil
	c.UpdateDimensions()
	if c.Rows == 0 || c.Cols == 0 {
		err = fmt.Errorf("the console dimensions could not be initialized")
	}
	return
}

// UpdateDimensions updates the ConsoleHelper instance with the current console dimensions.
func (c *ConsoleHelper) UpdateDimensions() {
	c.Cols, c.Rows = consolesize.GetConsoleSize()
}

// _updateDimensions updates the ConsoleHelper instance with the current console dimensions if necessary.
func (c *ConsoleHelper) _updateDimensions() {
	if (c.Cols == 0 || c.Rows == 0) || c.AutoUpdate {
		c.UpdateDimensions()
	}
}

// Fprintf is like fmt.Fprintf, but with text wrapping based on console size.
func (c *ConsoleHelper) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c._updateDimensions()
	return fmt.Fprintf(w, c.Sprintf(format, a...))
}

// Fprint is like fmt.Fprint, but with text wrapping based on console size.
func (c *ConsoleHelper) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c._updateDimensions()
	return fmt.Fprint(w, c.Sprint(a...))
}

// Fprintln is like fmt.Fprintln, but with text wrapping based on console size.
func (c *ConsoleHelper) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	c._updateDimensions()
	return fmt.Fprintln(w, c.Sprint(a...))
}

// Printf is like fmt.Printf, but with text wrapping based on console size.
func (c *ConsoleHelper) Printf(format string, a ...interface{}) (n int, err error) {
	c._updateDimensions()
	return fmt.Printf(c.Sprintf(format, a...))
}

// Print is like fmt.Print, but with text wrapping based on console size.
func (c *ConsoleHelper) Print(a ...interface{}) (n int, err error) {
	c._updateDimensions()
	return fmt.Print(c.Sprint(a...))
}

// Println is like fmt.Println, but with text wrapping based on console size.
func (c *ConsoleHelper) Println(a ...interface{}) (n int, err error) {
	c._updateDimensions()
	return fmt.Println(c.Sprint(a...))
}

// Sprintf is like fmt.Sprintf, but with text wrapping based on console size.
func (c *ConsoleHelper) Sprintf(format string, a ...interface{}) string {
	c._updateDimensions()
	return text.WrapSoft(fmt.Sprintf(format, a...), c.Cols)
}

// Sprint is like fmt.Sprint, but with text wrapping based on console size.
func (c *ConsoleHelper) Sprint(a ...interface{}) string {
	c._updateDimensions()
	return text.WrapSoft(fmt.Sprint(a...), c.Cols)
}

// Sprintln is like fmt.Sprintln, but with text wrapping based on console size.
func (c *ConsoleHelper) Sprintln(a ...interface{}) string {
	c._updateDimensions()
	//return text.WrapSoft(fmt.Sprintln(a...), c.Cols) // Does not work - final newline missing
	return fmt.Sprintln(text.WrapSoft(fmt.Sprint(a...), c.Cols))
}
