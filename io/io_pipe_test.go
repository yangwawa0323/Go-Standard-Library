// Golang program to illustrate the usage of
// io.Pipe() function

// Including main package
package io_pipe_test

// Importing fmt, io, and bytes
import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// Calling main
func Test_IO_Pipe(t *testing.T) {

	// Calling Pipe method
	pipeReader, pipeWriter := io.Pipe()

	// Using Fprint in go function to write
	// data to the file
	go func() {
		fmt.Fprint(pipeWriter, "GeeksforGeeks\nis\na\nCS-Portal.\n")

		// Using Close method to close write
		pipeWriter.Close()
	}()

	// Creating a buffer
	buffer := new(bytes.Buffer)

	// Calling ReadFrom method and writing
	// data into buffer
	buffer.ReadFrom(pipeReader)

	// Prints the data in buffer
	fmt.Printf("buffer :%q\n", buffer.String())
}

func Test_IO_Pipe2(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "Some io.Reader stream to be read\n")
		w.Close()
	}()

	// if _, err := io.Copy(os.Stdout, r); err != nil {
	// 	t.Fatal(err)
	// }

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		t.Fatal(err)
	}
	t.Log("buffer: ", buf.String())
}
