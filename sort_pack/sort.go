package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	Age  int
}

type people []person

func (p *people) Len() int {
	return len(*p)
}

func (p *people) Swap(f, s int) {
	temp := (*p)[f]
	(*p)[f] = (*p)[s]
	(*p)[s] = temp
}

func (p *people) Less(f, s int) bool {
	return (*p)[f].Age < (*p)[s].Age
}

func main() {

	dummyIntegersForBBSort := []int{5, 2, 3, 1, 4}
	bubbleSortForInts(dummyIntegersForBBSort)
	fmt.Println(dummyIntegersForBBSort)

	dummyIntegersForBuiltInSort := []int{50, 20, 30, 10, 40}
	sort.Ints(dummyIntegersForBuiltInSort)
	fmt.Println(dummyIntegersForBuiltInSort)

	peopleBBSort := []person{{name: "Stefan", Age: 25}, {name: "Meri", Age: 21}}
	bubbleSortForPersonStructs(peopleBBSort)
	fmt.Println(peopleBBSort)

	pslice := people([]person{{name: "Stefan", Age: 25}, {name: "Meri", Age: 21}})
	sort.Sort(&pslice)
	fmt.Println(pslice)

	peopleSliceSort := []person{{name: "Stefan", Age: 25}, {name: "Meri", Age: 21}}
	sort.Slice(peopleSliceSort, func(i, j int) bool { return peopleSliceSort[i].Age < peopleSliceSort[j].Age })
	fmt.Println(peopleSliceSort)

}

func bubbleSortForInts(dummyIntegers []int) []int {
	for i := 0; i < len(dummyIntegers)-1; i++ {
		for j := 0; j < len(dummyIntegers)-i-1; j++ {
			if dummyIntegers[j] > dummyIntegers[j+1] {
				swap := dummyIntegers[j]
				dummyIntegers[j] = dummyIntegers[j+1]
				dummyIntegers[j+1] = swap
			}
		}
	}
	return dummyIntegers
}

func bubbleSortForPersonStructs(dummyStructs []person) []person {
	for i := 0; i < len(dummyStructs)-1; i++ {
		for j := 0; j < len(dummyStructs)-i-1; j++ {
			if dummyStructs[j].Age > dummyStructs[j+1].Age {

				swap := dummyStructs[j]
				dummyStructs[j] = dummyStructs[j+1]
				dummyStructs[j+1] = swap

			}
		}
	}
	return dummyStructs
}
