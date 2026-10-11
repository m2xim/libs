package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteFileNestedDir(t *testing.T) {
	dir := t.TempDir()
	e := newExtractor(dir, nil)
	if err := e.writeFile("a/b/c.txt", strings.NewReader("hello"), 5); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "a", "b", "c.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello" {
		t.Fatalf("unexpected content: %q", data)
	}
}
