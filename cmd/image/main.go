package main

import (
	"bytes"
	"fmt"

	"github.com/trystanj/picam/pkg/image"
)

func main() {
	output, err := image.Capture(1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	result := bytes.NewBuffer(output).String()
	fmt.Printf("Result: %s", result)
}
