package main

import (
	compo "designpatternwithgo/Structural/composite/composite_example"
)

func main(){
	file1 := &compo.File{Name: "File1"}
	file2 := &compo.File{Name: "File2"}
	file3 := &compo.File{Name: "File3"}
	file4 := &compo.File{Name: "File4"}
	folder1 := &compo.Folder{
		Name: "Folder1",
	}
	folder1.Add(file1)
	folder2 := &compo.Folder{
		Name: "Folder2",
	}
	folder3 := &compo.Folder{
		Name: "Folder3",
	}
	folder3.Add(file4)
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)
	folder2.Add(folder3)
	folder2.Search("File2")
}
