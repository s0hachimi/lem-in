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
	mp     = make(map[string][]string)
	slofsl [][]string
)

func main() {
	Readfile()
	LinksFinder(Ants.Links)
	FindPath()
}

func FindPath() {
	var Path []string

	check := make(map[string]bool)

	// var Paths [][]string

	for i := 0; i <= len(slofsl)-1; i++ {
		for j := 0; j <= len(slofsl[i])-1; j++ {
			if len(Path) == 0 {
				Path = append(Path, slofsl[i][0])
				check[slofsl[i][j]] = true
				Path = append(Path, slofsl[i][j+1])
			}

		
			
			if !check[slofsl[i][j]] {
				if Path[len(Path)-1] == slofsl[i][0] {
					Path = append(Path, slofsl[i][j+1])
					check[slofsl[i][j]] = true
					check[slofsl[i][j+1]] = true
					i = 0 
					break
				}
			}
			

		}
	}
	fmt.Println(Path)
}

func LinksFinder(Links [][]string) {
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

	sl := []string{}

	for room, links := range mp {
		sl = append(sl, room)
		sl = append(sl, links...)
		slofsl = append(slofsl, sl)
		sl = []string{}
	}
	// fmt.Println(mp)
	fmt.Println(slofsl)
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

// for i := 0; i <= len(slofsl)-1; i++ {
// 	again:
// 		for j := 0; j <= len(slofsl[i])-1; j++ {
// 			if len(Path) != 0 {
// 				for _, p := range Path {
// 					if p != slofsl[i][j] {
// 						if slofsl[i][j] == slofsl[i][0] {
// 							Path = append(Path, slofsl[i][0])
// 							Path = append(Path, slofsl[i][j+1])
// 						}
// 						if slofsl[i][0] == Path[len(Path)-1] {
// 							if p != slofsl[i][j] {
// 								Path = append(Path, slofsl[i][j])
// 								i = 0
// 								j = 0
// 								goto again
// 							}
// 						}
// 					}
// 				}
// 			}
// 			Path = append(Path, slofsl[i][0])
// 			Path = append(Path, slofsl[i][j])
