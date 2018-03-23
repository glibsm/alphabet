package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var sampleData = []byte("This is some sample data in a file")

func main() {
	forAlphabet(func(l1 string) {
		forAlphabet(func(l2 string) {
			dirP := filepath.Join(l1, l2) // a/a a/b a/c ...
			err := os.MkdirAll(dirP, os.ModePerm)
			if err != nil {
				log.Panic(err)
			}

			forAlphabet(func(l3 string) {
				fileP := filepath.Join(dirP, l3) // a/a/a a/a/b /a/a/c ...

				err := ioutil.WriteFile(fileP, sampleData, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}
			})
		})
	})
}

func forAlphabet(f func(letter string)) {
	for _, b := range alphabet() {
		f(string(b))
	}
}

func alphabet() []byte {
	var err error
	b := &bytes.Buffer{}
	for i := 0; i < 26; i += 1 {
		err = b.WriteByte('a' + byte(i))
		if err != nil {
			log.Panic(err)
		}
	}
	return b.Bytes()
}
