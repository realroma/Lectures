package filer

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	want := map[string]string{
		"Some": "value",
	}
	got := ReadFile()
	if got["Some"] != want["Some"] {
		t.Fatalf("Получили %v, ожидалось %v", got["Some"], want["Some"])
	}
}

func TestWriteFile(t *testing.T) {
	want := `{"Some":"value"}`
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
