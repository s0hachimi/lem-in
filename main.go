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
	mp = make(map[string][]string)
	// slofsl [][]string
	Paths [][]string
	Path  []string
	check = make(map[string]bool)
)

func main() {
	Readfile()
	Graph := LinksFinder(Ants.Links)
	FindPath(Graph, Ants.Start, Ants.End, Path, check, &Paths)
}

func FindPath(Graph map[string][]string, start, end string, Path []string, check map[string]bool, Paths *[][]string) {
	check[start] = true
	Path = append(Path, start)

	if start == end {
		pathCopy := make([]string, len(Path))
		copy(pathCopy, Path)
		*Paths = append(*Paths, pathCopy)
	} else {
		for _, n := range Graph[start] {
			if !check[n] {
				FindPath(Graph, n, Ants.End, Path, check, Paths)
			}
		}
	}

	if len(Path) > 1 {
		Path = Path[:len(Path)-1]
		check[start] = false
	}

	fmt.Println(Paths)
}

func LinksFinder(Links [][]string) map[string][]string {
	for _, sl := range Links {
		for _, str := range sl {
			_, ok := mp[str]
			if ok {
				continue
			}
			for _, sl := range Links {
				for i, str2 := range sl {
					if str == str2 {
						if i == 0 {
							mp[str] = append(mp[str], sl[i+1])
						} else {
							mp[str] = append(mp[str], sl[i-1])
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
	// fmt.Println(mp)
	// fmt.Println(slofsl)

	return mp
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
