package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// rot13 applies the ROT13 cipher to a byte slice.
func rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case b >= 'a' && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b // Return unchanged if not an alphabet character
	}
}

// Read reads from the underlying io.Reader and applies the ROT13 transformation.
func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p) // Read from the underlying reader
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i]) // Apply ROT13 to each byte
	}
	return n, err // Return the number of bytes read and any error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // Input string
	r := rot13Reader{s} // Wrap the string reader with rot13Reader
	io.Copy(os.Stdout, &r) // Copy the transformed output to stdout
}
