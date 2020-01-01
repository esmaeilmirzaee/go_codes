package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}

	if !strings.Contains(d.toString(), "One of Diamonds") {
		t.Errorf("Unvalid cards")
	}
}

func TestSaveDackToFileAndLoadDeckFromFile(t *testing.T) {
	os.Remove("_testingdeck")

	deck := newDeck()
	deck.saveToFile("_testingdeck")

	loadedDeck := newDeckFromFile("_testingdeck")

	if len(loadedDeck) != 520 {
		t.Errorf("Not loading deck correctly")
	}

	os.Remove("_testingdeck")
}
