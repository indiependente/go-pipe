package main

import (
	"io"
	"os"

	"github.com/indiependente/pipe/transform"
)

func main() {
	echo := &transform.Echo{}
	io.Copy(os.Stdout, echo.Transform(os.Stdin))
}

func getDefault(variable string, defalut string) string {
	value := os.Getenv(variable)
	if value == "" {
		return defalut
	}
	return value
}
