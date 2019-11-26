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
