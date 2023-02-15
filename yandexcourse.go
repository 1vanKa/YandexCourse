package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type limitReader struct {
	reader io.Reader
	left   int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.left == 0 {
		return 0, io.EOF
	}
	if l.left < len(p) {
		p = p[:l.left]
	}
	n, err = l.reader.Read(p)
	l.left -= n
	return n, err
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{
		reader: r,
		left:   n,
	}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
