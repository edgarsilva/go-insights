package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	str := "Hello World of Golang!"
	src := StrReader{str: &str}
	dst := &StrWriter{}

	printSrcToDst(src, dst)

	fmt.Println("Full written dst:", string(dst.data))
}

type StrReader struct {
	str *string
}

func (s StrReader) Read(buf []byte) (int, error) {
	if len(*s.str) == 0 {
		return 0, io.EOF
	}

	n := copy(buf, *s.str)
	*s.str = (*s.str)[n:]

	return n, nil
}

type StrWriter struct {
	data []byte
}

func (w *StrWriter) Write(buf []byte) (int, error) {
	fmt.Println("writing ->", string(buf))
	w.data = append(w.data, buf...)
	return len(buf), nil
}

func printSrcToDst(src io.Reader, dst io.Writer) {
	buf := make([]byte, 8)

	for {
		n, err := src.Read(buf)
		if n > 0 {
			_, err := dst.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}
