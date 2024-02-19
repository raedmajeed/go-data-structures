package pkg

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SortMethods() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}
	sort.Sort(ByAge(people))
	fmt.Println(people) // Output: [{Bob 20} {Alice 25} {Charlie 30}]
}
