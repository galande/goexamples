package main

import (
	"fmt"
	"sort"
)

func main() {

	i := []int{1, 2, 3, 8, 4, 6, 5}

	//sort.Ints()

	fmt.Printf("%T \n", sort.IntSlice(i))

	is := sort.Reverse(sort.IntSlice(i))
	sort.Sort(is)
	fmt.Printf("%T \n", is)
	fmt.Println(is)

	// sort.Sort(sort.Reverse(sort.IntSlice(i)))
	// fmt.Println(i)
}
