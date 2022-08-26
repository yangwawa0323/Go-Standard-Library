package os_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func capture(f func()) (string, error) {
	// Keep the original os.Stdout in case to restore back.
	rescueOut := os.Stdout
	// Make a Pipe with two *File, one for reader , other for writer.
	// anything writing to writer would reading from ther reader.
	r, w, err := os.Pipe()

	if err != nil {
		return "", fmt.Errorf("capture: %s", err.Error())
	}
	// TeeReader copy the `r`(source reader) to `rescueOut`(destination writer),
	// and return a new Reader
	tee := io.TeeReader(r, rescueOut)

	// the standard output is replaced with pipe's writer, any output will write to pipe's writer,
	// and at other side you can reads the info from pipe's reader.
	//  message ---> ( os.Stdout ---> pipe.Writer ) --->  (os.Pipe | ) ---> pipe.Reader ---> message
	os.Stdout = w

	outC := make(chan string)

	go func() {
		// out, err := ioutil.ReadAll(r)
		var buf bytes.Buffer
		_, err := io.Copy(&buf, tee) // anything came from tee will be copy to buf

		if err != nil {
			log.Fatalf("capture: %s", err.Error())
		}
		// outC <- string(out)
		outC <- buf.String()
	}()

	// this f() produces the output message to os.Stdout
	f()

	if err := w.Close(); err != nil {
		return "", fmt.Errorf("capture: %s", err.Error())
	}

	// restore the original os.Stdout back.
	os.Stdout = rescueOut

	return <-outC, nil
}

func Test_Capture_Func_Output(t *testing.T) {
	output, _ := capture(func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Fprint(os.Stdout, time.Now().String()+"\n")
		}
	})

	t.Log("================================")
	t.Log("Final output: \n", output)
}

func Test_OS_Pipe(t *testing.T) {

	rd, wr, _ := os.Pipe()

	defer func() {
		wr.Close()
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			wr.Write([]byte(time.Now().String()))
		}
	}()

	for {
		var dataRead = make([]byte, 256)
		n, _ := rd.Read(dataRead)
		fmt.Println(string(dataRead[:n]))
	}
}

func Test_IO_Pipe(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		writer := bufio.NewWriter(w)
		for {
			time.Sleep(1 * time.Second)
			writer.WriteString(time.Now().String() + "\n")
			// buffer size is 4096 by default, there is no output utils the buffer is full
			// You need flush the buffer
			writer.Flush()
		}
	}()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
