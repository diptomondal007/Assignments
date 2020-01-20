package composite_example

import "fmt"

type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}

func (f *File) getName() string {
	return f.Name
}