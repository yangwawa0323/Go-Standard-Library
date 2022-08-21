// Golang program to illustrate the usage of
// io.Pipe() function

// Including main package
package io_pipe

// Importing fmt, io, and bytes
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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

func Test_IO_WriteString(t *testing.T) {
	if _, err := io.WriteString(os.Stdout, "Hello world"); err != nil {
		t.Fatal(err)
	}
}

func Test_LimitedReader(t *testing.T) {
	type myLimitedReader struct {
		lmtRead io.LimitedReader
	}

	buf := new(bytes.Buffer)
	buf.WriteString("Hello world!")

	lmt_read := myLimitedReader{
		lmtRead: io.LimitedReader{R: buf, N: 3},
	}

	n, err := io.Copy(os.Stdout, &lmt_read.lmtRead)
	if err != nil {
		log.Fatal("io.Copy: ", err)
	}

	t.Logf("limited reader op: %d\n", n)

}
