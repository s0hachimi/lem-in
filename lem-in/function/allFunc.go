package function

import (
	"fmt"
)

type Ant struct {
	count int
	Start string
	End   string
	Rooms []string
	Links [][]string
}

var Ants Ant

var (
	Graph   = make(map[string][]string)
	Paths   [][]string
	Groupes = [][][]string{}
)

func Sort(Paths [][]string) {
	for i := 0; i <= len(Paths)-1; i++ {
		for j := 0; j <= len(Paths)-1; j++ {
			if len(Paths[i]) < len(Paths[j]) {
				Paths[i], Paths[j] = Paths[j], Paths[i]
			}
		}
	}
}

func FindPath(start, end string) {
	visited := make(map[string]bool)
	Path := []string{}
	dfs(start, end, visited, Path)
}

func dfs(start, end string, visited map[string]bool, Path []string) {
	visited[start] = true
	Path = append(Path, start)

	if start == end {

		pathCopy := make([]string, len(Path))
		copy(pathCopy, Path)
		Paths = append(Paths, pathCopy)

	} else {
		for _, neighbor := range Graph[start] {
			if !visited[neighbor] {
				dfs(neighbor, Ants.End, visited, Path)
			}
		}
	}
	visited[start] = false
}

// func FilterPath(Paths [][]string) {

// }

func FilterPath(Paths [][]string) {
	Groupe := [][]string{}

	for i := 0; i <= len(Paths)-1; i++ {
		check := make(map[string]bool)
		for j := 1; j <= len(Paths[i])-2; j++ {
			check[Paths[i][j]] = true
		}

		Groupe = append(Groupe, Paths[i])

		for i := 0; i <= len(Paths)-1; i++ {
			if allfalse(Paths[i], check) {

				Groupe = append(Groupe, Paths[i])
				for k := 1; k <= len(Paths[i])-2; k++ {
					check[Paths[i][k]] = true
				}
			}
		}
		Groupes = append(Groupes, Groupe)
		Groupe = [][]string{}

	}

	// for i := 0; i <= len(Groupes)-1; i++ {
	// 	for j := 0; j <= len(Groupes)-1; j++ {
	// 		if len(Groupes[i]) > len(Groupes[j]) {
	// 			Groupes[i], Groupes[j] = Groupes[j], Groupes[i]
	// 		}
	// 	}
	// }

	fmt.Println("\n\n", Groupes)
}

// func choosePath(count int) map[int][][][]string {
// 	fmt.Println(count)
// 	scoree := 0
// 	score := make(map[int][][][]string)
// 	for i := 0; i <= len(Groupes)-1; i++ {
// 		for j := 0; j <= len(Groupes[i])-1; j++ {
// 			for k := 1; k <= len(Groupes[i][j])-2; k++ {
// 				scoree++
// 			}
// 		}
// 		scoree = scoree + (count - len(Groupes[i]))

// 		score[scoree] = append(score[scoree], Groupes[i])
// 		scoree = 0
// 	}
// 	// fmt.Println(score)
// 	return score
// }

func allfalse(s []string, check map[string]bool) bool {
	for i := 1; i <= len(s)-2; i++ {
		if check[s[i]] {
			return false
		}
	}
	return true
}

func LinksFinder(Links [][]string) {
	for _, sl := range Links {
		for _, str := range sl {
			_, ok := Graph[str]
			if ok {
				continue
			}
			for _, sl1 := range Links {
				for i, str2 := range sl1 {
					if str == str2 {
						if i == 0 {
							Graph[str] = append(Graph[str], sl1[i+1])
						} else {
							Graph[str] = append(Graph[str], sl1[i-1])
						}
					}
				}
			}
		}
	}

	// sl := []string{}

	// for room, links := range mp {
	// 	sl = append(sl, room)
	// 	sl = append(sl, links...)
	// 	slofsl = append(slofsl, sl)
	// 	sl = []string{}
	// }
	fmt.Println(Graph)
	// fmt.Println(slofsl)
}
