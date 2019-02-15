package main

import (
	"fmt"
)

func printNames(n map[string]string) {
	for number, name := range n {
		fmt.Println(number, "is", name)
	}
}

func printNotes(n []string) {
	for _, n := range n {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
func mapNotes(notes []string, start int, step int) []string {
	copy := make([]string, len(notes))
	for i := range notes {
		copy[i] = notes[(start+step*i)%len(notes)]
	}
	return copy
}

var modeNames = map[string]string{
	"I":   "Ionian",
	"II":  "Dorian",
	"III": "Phrygian",
	"IV":  "Lydian",
	"V":   "Mixolydian",
	"VI":  "Aeolian",
	"VII": "Locrian",
}

var chordNames = map[string]string{
	"I":   "maj7",
	"II":  "m7",
	"III": "m7",
	"IV":  "maj7",
	"V":   "7",
	"VI":  "m7",
	"VII": "m7b5",
}

var cMajScale = []string{"C", "D", "E", "F", "G", "A", "B"}
var flatKeys = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "Cb"}
var flatNotes = mapNotes(cMajScale, 6, 3)
var sharpKeys = []string{"G", "D", "A", "E", "B", "F#", "C#"}
var sharpNotes = mapNotes(cMajScale, 3, 4)
