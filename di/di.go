package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
  fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
  Greet(os.Stdout, "Elodi")
}
