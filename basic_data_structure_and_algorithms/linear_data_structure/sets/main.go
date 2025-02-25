package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func NewSet() *Set {
	return &Set{
		integerMap: make(map[int]bool),
	}
}

func (set *Set) ContainsElement(element int) bool {
	_, exist := set.integerMap[element]

	return exist
}

func (set *Set) AddElement(element int) {
	set.integerMap[element] = true
}

func (set *Set) DeleteElement(value int) {
	delete(set.integerMap, value)
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	intersect := NewSet()

	for k := range set.integerMap {
		if anotherSet.ContainsElement(k) {
			intersect.AddElement(k)
		}
	}

	return intersect
}

func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := NewSet()

	for k := range set.integerMap {
		unionSet.AddElement(k)
	}

	for k := range anotherSet.integerMap {
		unionSet.AddElement(k)
	}

	return unionSet
}

func main() {
	set := NewSet()
	set.AddElement(1)
	set.AddElement(2)

	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))

	anotherSet := NewSet()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println(set.Intersect(anotherSet))
	fmt.Println(set.Union(anotherSet))
}
