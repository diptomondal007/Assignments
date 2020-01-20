package composite_example

import "fmt"

type Folder struct {
	components []Component
	Name string
}

func (f *Folder) Search(keyword string){
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.components{
		composite.Search(keyword)
	}
}

func(f *Folder) Add(c Component){
	f.components = append(f.components, c)
}
