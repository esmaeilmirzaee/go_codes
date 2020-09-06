package main

import (
	"math/rand"
	"strings"
	"time"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// another solution by me 😍
func testShuffle() []string {
	rand.Seed(time.Now().UnixNano())
	words := strings.Fields("Esmaeil MIRZAEE does not work really hard.")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return words
}
