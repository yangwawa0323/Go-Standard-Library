package writer

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_OS_OpenFile(t *testing.T) {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	now_str := time.Now().Format(time.RFC1123)

	var buf = bytes.NewBuffer([]byte(fmt.Sprintf("%s \n", now_str)))
	f.Write(buf.Bytes())
}
