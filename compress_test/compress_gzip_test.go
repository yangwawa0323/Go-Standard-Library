package compress_test

import (
	"compress/gzip"
	"os"
	"testing"
	"time"
)

func Test_Compress_GZip_Writer(t *testing.T) {

	file, err := os.Create("test.gz")
	if err != nil {
		t.Fatal("Can not create the file")
	}

	zw := gzip.NewWriter(file)

	// Setting the Header fileds is optional
	zw.Name = "a-new-compress.text"
	zw.Comment = "lorem"
	zw.ModTime = time.Date(1997, time.July, 19, 0, 0, 0, 0, time.UTC)

	_, err = zw.Write([]byte("A long time ago in a galaxy far, far away.."))
	if err != nil {
		t.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
}

func Test_Compress_Gzip_Multiple_File(t *testing.T) {
	file, err := os.Create("mutilple_files.gz")
	if err != nil {
		t.Fatal(err)
	}

	zw := gzip.NewWriter(file)

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Date(2007, time.March, 2, 4, 5, 6, 1, time.UTC), "Hello Gophers - 2"},
	}

	for _, f := range files {
		zw.Name = f.name
		zw.Comment = f.comment
		zw.ModTime = f.modTime
		if _, err := zw.Write([]byte(f.data)); err != nil {
			t.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			t.Fatal(err)
		}

		zw.Reset(file)
	}
}
