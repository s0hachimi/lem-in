package main

import (
	"fmt"
	"lem-in/function"
)


func main() {
	function.Readfile()
	function.LinksFinder(Ants.Links)
	function.FindPath(Ants.Start, Ants.End)
	function.Sort(Paths)

	fmt.Println(Paths)
	function.FilterPath(Paths)
	// mm := choosePath(Ants.count)
	// fmt.Println("\n\n", mm)
}

// func Sort(Paths [][]string) {
// 	sort.Slice(Paths, func(i, j int) bool {
// 		return len(Paths[i]) < len(Paths[j])
// 	})
// }
