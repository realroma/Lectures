package filer

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	want := map[string]string{
		"Some": "value",
	}

	f := New("Link.txt")
	got := Read(f)

	if got["Some"] != want["Some"] {
		t.Fatalf("Получили %v, ожидалось %v", got["Some"], want["Some"])
	}
}

func TestWriteFile(t *testing.T) {
	want := `{"Some":"value"}`
	m := map[string]string{
		"Some": "value",
	}

	am, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	f := New("Link.txt")
	Write(f, am)

	text, err := os.ReadFile("Link.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := string(text)
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}
