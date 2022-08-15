package bytes_buffer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"sort"
	"testing"
	"unicode"
)

func Test_BytesBuffer(t *testing.T) {
	// A Buffer needs no initialization
	var b bytes.Buffer

	// Write appends the contents to the buffer
	// growing the buffer as needed.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}

func Test_NewBufferString_Decoder(t *testing.T) {
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

func Test_BufferRead(t *testing.T) {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	rdbuf := make([]byte, 2)
	n, err := b.Read(rdbuf)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(n)          // How many number of bytes has been readed
	fmt.Println(b.String()) // Remaining bytes in the buffer
	fmt.Println(string(rdbuf))
}

func Test_BufferReadByte(t *testing.T) {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	c, err := b.ReadByte()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(c)          // the byte has been readed to `c`
	fmt.Println(b.String()) // Remaining bytes in the buffer
}

func Test_BufferNext(t *testing.T) {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%s\n", string(b.Next(2)))
	fmt.Printf("%s\n", string(b.Next(2)))
	fmt.Printf("%s\n", string(b.Next(2)))
}

func Test_BufferCompare(t *testing.T) {
	// var a , b []byte = []byte("abc") , []byte("bcd")
	var needle []byte = []byte("foo")
	var haystack [][]byte = [][]byte{[]byte("foo"), []byte("bar")} // Assume sorted
	i := sort.Search(len(haystack), func(i int) bool {
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		fmt.Println("Found it")
	}
}

func Test_BytesFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	var testString []byte = []byte("  foo1;bar2,baz3...")
	fmt.Printf("Fields are: %q", bytes.FieldsFunc(testString, f))
}

func Test_UnicodeHan(t *testing.T) {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f))
}
