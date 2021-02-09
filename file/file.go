package file

import (
	"io/ioutil"
	"log"
)

// NewFile reads a file and returns a File object
func NewFile(name string) *File {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	return &File{name: name, value: text}
}

// File for evaluating programs from files
type File struct {
	name  string
	value string
}

func (f *File) String() string {
	return f.value
}
