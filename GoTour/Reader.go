package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		//You can either use 65 or 'A'
		b[i] = 65
	}
	return len(b), nil
}
func main() {
	reader.Validate(MyReader{})
}
