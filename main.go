package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Ants struct {
	count int
	Start string
	End   string
	Rooms []string
	Links [][]string
}

var (
	Graph = make(map[string][]string)
	Paths [][]string
)

func main() {
	Readfile()
	LinksFinder(Ants.Links)
	FindPath(Ants.Start, Ants.End)
	fmt.Println(Paths)
	FilterPath(Paths)
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
	// Path = Path[:len(Path)-1]
	visited[start] = false
}

func FilterPath(Paths [][]string) {
	Groupes := [][][]string{}

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

	fmt.Println(Groupes)

}

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

func Readfile() {
	text, err := os.ReadFile("test01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(text), "\n")
	n, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	Ants.count = n

	for i := 0; i <= len(lines)-1; i++ {

		if lines[i] == "" || len(lines[i]) != 0 && string(lines[i][0]) == "#" && string(lines[i][1]) != "#" {
			continue
		}

		if lines[i] == "##start" {
			for j := 0; j <= len(lines[i+1])-1; j++ {
				if lines[i+1][j] == ' ' {
					Ants.Start = string(lines[i+1][:j])
					break
				}
			}
		}

		if lines[i] == "##end" {
			for j := 0; j <= len(lines[i+1])-1; j++ {
				if lines[i+1][j] == ' ' {
					Ants.End = string(lines[i+1][:j])
					break
				}
			}
		}

		if strings.Contains(lines[i], "-") {
			Ants.Links = append(Ants.Links, strings.Split(lines[i], "-"))
		} else {
			for j := 0; j <= len(lines[i])-1; j++ {
				if lines[i][j] == ' ' {
					Ants.Rooms = append(Ants.Rooms, string(lines[i][:j]))
					break
				}
			}
		}

	}

	// fmt.Println(Ants.Nmala)
	// fmt.Println(Ants.Start)
	// fmt.Println(Ants.End)
	// fmt.Println(Ants.Links)
	// fmt.Printf("%#v\n",Ants.Rooms)
	// fmt.Printf("%#v\n",Ants)
	// fmt.Printf("%#v\n", sl)
}
