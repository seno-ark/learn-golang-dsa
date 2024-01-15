package main

import (
	"fmt"
)

// Set class
type Set struct {
	intMap map[int]bool
}

func (set *Set) New() {
	set.intMap = make(map[int]bool)
}

// containsElement checks if element is in the set
func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.intMap[element]

	return exists
}

func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.intMap[element] = true
	}
}

// DeleteElement deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.intMap, element)
}

// Intersect returns a new set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()

	var key int
	for key = range set.intMap {
		if anotherSet.ContainsElement(key) {
			intersectSet.AddElement(key)
		}
	}
	return intersectSet
}

// Union returns the set which is union of the set with anotherSet
func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()

	var key int
	for key = range set.intMap {
		unionSet.AddElement(key)
	}

	for key = range anotherSet.intMap {
		unionSet.AddElement(key)
	}

	return unionSet
}

func main() {
	set1 := &Set{}
	set1.New()

	set1.AddElement(1) // 1
	set1.AddElement(2) // 1,2

	fmt.Println("initial set1", set1)    // 1, 2
	fmt.Println(set1.ContainsElement(1)) // true

	set2 := &Set{}
	set2.New()

	set2.AddElement(2) // 2
	set2.AddElement(4) // 2,4
	set2.AddElement(5) // 2,4,5

	fmt.Println("another set2", set2) // 2, 4, 5

	fmt.Println("intersection of sets ", set1.Intersect(set2)) // 2

	fmt.Println("union of sets ", set1.Union(set2)) // 1, 2, 4, 5
}
