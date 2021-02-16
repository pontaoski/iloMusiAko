package main

import "math/rand"

var letterData = []string{"k", "l", "m", "n", "p", "s", "t", "w", "j", "a", "e", "i", "o", "u", "k", "l", "m", "n", "p", "s", "t", "w", "j", "a", "e", "i", "o", "u", "k", "l", "m", "n", "p", "s", "t", "w", "j", "a", "e", "i", "o", "u"}
var particles = map[string]struct{}{
	"li": {},
	"e":  {},
	"la": {},
	"pi": {},
}

func randomLetter() string {
	return letterData[rand.Intn(len(letterData))]
}

func randomLength() int {
	return rand.Intn(3) + 2
}

func randomLetters() []string {
	length := randomLength()
	arr := make([]string, length)
	for i := 0; i < length; i++ {
		arr[i] = randomLetter()
	}
	return arr
}
