/*

Package consolehelper aids with pretty printing on the CLI.

	package main

	import (
		"fmt"
		consolehelper "gitlab.com/rbrt-weiler/go-module-consolehelper"
	)

	func main() {
		var cons comsolehelper.ConsoleHelper
		fmt.Printf("%s\n", cons.Sprint("This line will be wrapped according to the comsole width."))
	}

*/
package consolehelper

import (
	"fmt"

	text "github.com/jedib0t/go-pretty/v6/text"
	consolesize "github.com/nathan-fiscaletti/consolesize-go"
)

const (
	// ModuleName contains the short name of the module.
	ModuleName string = "consolehelper"
	// ModuleVersion contains the version number of the module.
	ModuleVersion string = "0.1.0"
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
	/*
		// Currently not feasible - consolesize will panic during testing
		// TODO: Revise once consolesize is test-safe.
		c.UpdateDimensions()
		if c.Rows == 0 || c.Cols == 0 {
			c.Rows = -1
			c.Cols = -1
			err = fmt.Errorf("the console dimensions could not be initialized")
		}
	*/
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
