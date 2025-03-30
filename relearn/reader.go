package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func reader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

//Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i, char := range b[:n] {
		if char >= 'A' && char <= 'Z' {
			b[i] = 'A' + (char-'A'+13)%26
		} else if char >= 'a' && char <= 'z' {
			b[i] = 'a' + (char-'a'+13)%26
		}
	}
	return n, err
}

func reader2() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
