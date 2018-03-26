package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	rand.Seed(time.Now().UnixNano())

	forAlphabet(func(l1 string) {
		forAlphabet(func(l2 string) {
			dirP := filepath.Join(l1, l2) // a/a a/b a/c ...
			err := os.MkdirAll(dirP, os.ModePerm)
			if err != nil {
				log.Panic(err)
			}

			forAlphabet(func(l3 string) {
				fileP := filepath.Join(dirP, l3) // a/a/a a/a/b /a/a/c ...

				err := ioutil.WriteFile(
					fileP,
					randomLetters(16), // write out 16 random letters
					os.ModePerm,
				)
				if err != nil {
					log.Fatal(err)
				}
			})
		})
	})
}

func forAlphabet(f func(letter string)) {
	for _, b := range alphabet {
		f(string(b))
	}
}

func randomLetters(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
