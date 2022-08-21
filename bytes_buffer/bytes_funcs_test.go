package bytes_buffer

import (
	"bytes"
	"testing"
)

func Test_TrimSuffix(t *testing.T) {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	t.Logf("%q\n", b)
	b = bytes.TrimSuffix(b, []byte("gopher"))
	t.Logf("%q\n", b)
	b = append(b, bytes.TrimSuffix([]byte("world"), []byte("x!"))...)
	t.Logf("%q\n", b)
}

func show(t *testing.T, s, sep string) {
	before, after, found := bytes.Cut([]byte(s), []byte(sep))
	if found {
		t.Logf("Cur(%q , %q) = %q, %q ,%v\n", s, sep, before, after, found)
		show(t, string(after), sep)
	}
	return
}
func Test_Cut(t *testing.T) {

	show(t, "Gopher Go Go Go", "Go")
	show(t, "Gopher Go Go Go", "ph")
	show(t, "Gopher Go Go Go", "er")
	show(t, "Gopher Go Go Go", "Badger")
}

func Test_Fields(t *testing.T) {
	t.Logf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}

func Test_Split(t *testing.T) {
	t.Logf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	t.Logf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	t.Logf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	t.Logf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}
