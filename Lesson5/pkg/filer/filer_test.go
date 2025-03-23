package filer

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	CheckSave()
}

func TestWriteFile(t *testing.T) {
	want := `{"Some":1}`
	WriteFile()
	text, err := os.ReadFile("Link.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := string(text)
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}
