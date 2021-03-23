package dealFunc

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

func Func1() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("could not open file %q: %v", os.Args[1], err)
	}

	words := 0
	inWord := false
	for {
		r, err := readByte1(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inWord {
			words++
			inWord = false
		}
		inWord = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}


func readByte1(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}