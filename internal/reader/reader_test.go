package reader

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("./testdata/input01.txt")
	if err != nil {
		t.Fatalf("failed to open file: %#v", err)
	}

	s := bufio.NewScanner(file)
	if s.Text() != "" {
		t.Fatalf("want: %v\ngot: %v", "", s.Text())
	}

	if !s.Scan() {
		t.Fatalf("unexpected scan result:\n\twant: true\n\tgot: false")
	}
	if s.Text() != "abc" {
		t.Fatalf("want: %v\ngot: %v", "abc", s.Text())
	}

	if !s.Scan() {
		t.Fatalf("unexpected scan result:\n\twant: true\n\tgot: false")
	}
	if s.Text() != "def" {
		t.Fatalf("want: %v\ngot: %v", "def", s.Text())
	}

	// if !s.Scan() {
	// 	t.Fatalf("unexpected scan result:\n\twant: true\n\tgot: false")
	// }
	// if s.Text() != "" {
	// 	t.Fatalf("want: %v\ngot: %v", "", s.Text())
	// }

	// NOTE: Last blank line is ignored.
	if s.Scan() {
		t.Fatalf("unexpected scan result:\n\twant: false\n\tgot: true")
	}
}

func TestReadFileWithoutTailingBlankLine(t *testing.T) {
	file, err := os.Open("./testdata/input02.txt")
	if err != nil {
		t.Fatalf("failed to open file: %#v", err)
	}

	s := bufio.NewScanner(file)
	if s.Text() != "" {
		t.Fatalf("want: %v\ngot: %v", "", s.Text())
	}

	if !s.Scan() {
		t.Fatalf("unexpected scan result:\n\twant: true\n\tgot: false")
	}
	if s.Text() != "abc" {
		t.Fatalf("want: %v\ngot: %v", "abc", s.Text())
	}

	if !s.Scan() {
		t.Fatalf("unexpected scan result:\n\twant: true\n\tgot: false")
	}
	if s.Text() != "def" {
		t.Fatalf("want: %v\ngot: %v", "def", s.Text())
	}

	// if !s.Scan() {
	// 	t.Fatalf("unexpected scan result:\n\twant: true\n\tgot: false")
	// }
	// if s.Text() != "" {
	// 	t.Fatalf("want: %v\ngot: %v", "", s.Text())
	// }

	// NOTE: Last blank line is ignored.
	if s.Scan() {
		t.Fatalf("unexpected scan result:\n\twant: false\n\tgot: true")
	}
}

func scannerFromFilename(filename string) *bufio.Scanner {
	f, err := os.Open("./testdata/input01.txt")
	if err != nil {
		log.Fatalf("failed to open file named %s: %v", filename, err)
	}
	return bufio.NewScanner(f)
}

func Test_charReader_NextChar(t *testing.T) {
	filename := "testdata/input01.txt"
	input, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open file named %s: %#v", filename, err)
	}
	charReader := NewCharReader(input, os.Stdout)

	// test
	wantChars := []rune{'a', 'b', 'c', ' ', 'd', 'e', 'f', ' '}
	for _, wantChar := range wantChars {
		gotChar, read, err := charReader.NextChar()
		if err != nil {
			t.Fatalf("failed to read char: %#v", err)
		}
		if !read {
			t.Fatal("EOF is not expected.")
		}
		if gotChar != wantChar {
			t.Fatalf("\nwant: %c\ngot: %c", wantChar, gotChar)
		}
	}

	wantread := false
	_, gotread, err := charReader.NextChar()
	if gotread != wantread {
		t.Fatalf("\nwant: %v\ngot: %v", wantread, gotread)
	}
}
