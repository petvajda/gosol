package main

import (
	"testing"
)

func TestNoteGeneration(t *testing.T) {
	if len(flatNotes) != 7 {
		t.Errorf("Expected flatNotes length of 7, but got %v", len(flatNotes))
	}

	if len(sharpNotes) != 7 {
		t.Errorf("Expected sharpNotes length of 7, but got %v", len(sharpNotes))
	}
	flats := []string{"B", "E", "A", "D", "G", "C", "F"}

	for i := range flats {
		if flatNotes[i] != flats[i] {
			t.Errorf("Expected flatnote %v, but got %v", flats[i], flatNotes[i])
		}
	}
	sharps := []string{"F", "C", "G", "D", "A", "E", "B"}
	for i := range sharps {
		if sharpNotes[i] != sharps[i] {
			t.Errorf("Expected sharpnote %v, but got %v", sharps[i], sharpNotes[i])
		}
	}

}
