# Go module consolehelper

consolehelper is a Go module that aids with pretty printing on the CLI. It provides [fmt](https://golang.org/pkg/fmt/)-like functionality with integrated text wrapping. The border for text wrapping is determined automatically, based on the console size.

## Example

```golang
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
```

## Documentation

### Module API

The API of this module is [fully documented](https://pkg.go.dev/gitlab.com/rbrt-weiler/go-module-consolehelper) on pkg.go.dev.

### Headless Environments

For headless environments a size of `0` will be returned for both console dimensions (rows and columns). This will effectively disable text wrapping.

### Variable Consoles

For consoles that are known or expected to change their size during the use of a program that imports consolehelper, auto-update functionality has been implemented. Set `AutoUpdate` to `true` to enable it.

```golang
cons, consErr := consolehelper.New()
cons.AutoUpdate = true
```

A drawback of the auto-update function is that it may result in `0` for one or both dimensions of the console. This will not result in an error or runtime panic, but will effectively disable text wrapping.

### Custom Console Size

The console size can be set manually by manipulating the `Cols` and `Rows` variables.

```golang
cons, consErr := consolehelper.New()
cons.Cols = 40
cons.Rows = 12
```

These settings will persist as long as `AutoUpdate` is disabled.

### Negative Console Dimensions

Using negative values for the console dimensions will have the following effects:

* `Cols < 0`: The string returned or printed will be empty.
* `Rows < 0`: No effect.

## Source

The original project is [hosted at GitLab](https://gitlab.com/rbrt-weiler/go-module-consolehelper), with a [copy over at GitHub](https://github.com/rbrt-weiler/go-module-consolehelper) for the folks over there.
