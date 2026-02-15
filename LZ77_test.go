package lz77

import (
	"fmt"
	"testing"
	"os"

	"github.com/urkx/lgo"
)

func TestLz(t *testing.T) {
	test := "tres tristes tigres tragaban trigo en un trigal"
	c := Compress(test, 32000)
	fmt.Print(lgo.Info("Compression finished"))
	fmt.Print(lgo.Debug(c))
	res := Decompress(c)

	if res != test {
		t.Fatal("Compression failed")
	}
}

func TestWriteFile(t *testing.T) {
	test := "tres tristes tigres tragaban trigo en un trigal"
	c := Compress(test, 32000)
	fmt.Print(lgo.Info("Compression finished"))
	fmt.Print(lgo.Info("Creating file"))
	err := WriteResultFile("test.lz77", c)

	if err != nil {
		t.Fatal("WriteFile failed")
	}
	fmt.Print(lgo.Info("File created successfully"))
}

func TestReadFile(t *testing.T) {
	test_content := "yipiyakei"
	err := os.WriteFile("test.input", []byte(test_content), 0644)
	if err != nil {
		t.Fatal("Could not write input test file", err)
	}

	read, error := ReadFile("test.input")
	if error != nil || string(read) != test_content {
		t.Fatal("Readed content is not equal to test content")
	}
}
