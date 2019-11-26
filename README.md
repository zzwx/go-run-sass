# go-run-sass

A simple command line tool wrapper to generate css using `go-libsass` which otherwise recompiles the whole library when simply imported as `github.com/wellington/go-libsass`.

## Installation

```
$ go get -u github.com/zzwx/go-run-sass
```

## Help

```
$ go-run-sass
Usage of go-run-sass:
  -help
        Print this help
  -i string
        Input SASS file
  -o string
        Output CSS file
```

## Usage in Code

```go
package main

import (
	"fmt"

	"github.com/zzwx/ifchanged"
)

func generateCSS() error {
	err := ifchanged.ExecuteCommand("go-run-sass", "-i", "./web/wireframe.scss", "-o", "./web/generated.css")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Done!\n")
	}
	return err
}

func main() {
	generateCSS()
}

```