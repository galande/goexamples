package main

import (
	"fmt"
	"sort"
)

func main() {

	i := []int{1, 2, 3, 8, 4, 6, 5}

	sort.Ints(i)

	fmt.Println(i)
}
