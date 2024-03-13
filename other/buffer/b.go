package main

import (
	"bytes"
	"fmt"
)

func cloneBuffer(src *bytes.Buffer) *bytes.Buffer {
	// Create a new destination buffer
	dest := new(bytes.Buffer)

	// Copy the content from the source buffer to the destination buffer
	dest.Write(src.Bytes())

	return dest
}

func main() {
	// Example source buffer
	sourceBuffer := bytes.NewBufferString("Hello, this is the source buffer.")

	// Clone the content into a new buffer
	destinationBuffer := cloneBuffer(sourceBuffer)

	// Modify the destination buffer without affecting the source buffer
	destinationBuffer.WriteString(" Additional content in the destination buffer.")

	// Print the original source buffer content
	fmt.Println("\nSource Buffer1:", sourceBuffer.String())
	fmt.Println("\nSource Buffer2:", sourceBuffer.String())

	// Print the modified destination buffer content
	fmt.Println("\nDestination Buffer:", destinationBuffer.String())
}
