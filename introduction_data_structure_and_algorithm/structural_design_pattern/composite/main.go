package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet: " + leaf.name)
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)

	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

func (branch *Branch) getLeafs() []Leaflet {
	return branch.leafs
}

func main() {
	branch := &Branch{
		name: "branch 1",
	}

	leaf1 := Leaflet{
		name: "leaf 1",
	}

	leaf2 := Leaflet{
		name: "leaf 2",
	}

	branch2 := Branch{
		name: "branch 2",
	}

	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)

	branch.perform()
}
