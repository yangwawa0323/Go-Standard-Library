package buffer_io

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func Test_Scanner(t *testing.T) {

	// An artificial input source.
	const input = "1234 5678  12345678901234567890"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int,
		token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		t.Logf("Invalid input: %s", err)
	}

	for _, word := range result {
		t.Logf("%s\n", word)
	}
}

func Test_Scaner_Split(t *testing.T) {
	// Comma-separated list; last entry is empty
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Define a split function
	onComma := func(data []byte, atEOF bool) (advance int,
		token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more token after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		t.Error("reading input:", err)
	}

	t.Logf("all: %q\n", result)

	for _, word := range result {
		t.Log(word)
	}

}

func Test_Scan_Line(t *testing.T) {

	// the artificial buffer string
	var buff_string = bytes.NewBuffer([]byte("This is first line\nThe Second line\nThe 3rd line."))

	scanner := bufio.NewScanner(buff_string)
	for scanner.Scan() {
		t.Log(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		t.Fatal("reading standard input: ", err)
	}
}

func Test_Scanner_Bytes(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("Hello world"))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t.Log(scanner.Text(), "length is equal 5 ", len(scanner.Bytes()) == 5)
	}

	if err := scanner.Err(); err != nil {
		t.Fatal("shouldn't see an error scanning a string")
	}

}
