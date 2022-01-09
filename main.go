package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var PNG_SIG = [8]byte{137, 80, 78, 71, 13, 10, 26, 10}
var filePath string

// Todo: Read file in chunks
func readBytes() []byte {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return b
}

func isPng(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func init() {
	flag.StringVar(&filePath, "fp", "", "")
}

func main() {
	flag.Parse()

	if len(filePath) == 0 {
		log.Fatal("Please provide a valid png file")
	}

	b := readBytes()
	var h []byte

	if len(b) < 8 {
		log.Fatal("File is not a png")
	} else {
		for i := 0; i <= 7; i++ {
			h = append(h, b[i])
		}
	}

	if isPng(PNG_SIG[:], h) {
		log.Println("Valid png")
	} else {
		log.Fatal("Not a png")
	}
}
