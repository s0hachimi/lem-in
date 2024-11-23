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
	Links string
}

var Graph struct {
	Path []string
}

func main() {
	Readfile()
	LinksFinder(Ants.Links)
}

func Paths() {
}

func LinksFinder(Links string) {
	
	

	roomAndLinks := make(map[string][]string)

	for i := 0; i <= len(Links)-1; i++ {
		
		if Links[i] == '*' {
			continue
		}

		for j := 0; j <= len(Links)-1; j++ {
			if Links[j] == '*' {
				continue
			}
			if Links[i] == Links[j] {
				if j != len(Links)-1 && Links[j+1] == '*' {
					roomAndLinks[string(Links[i])] = append(roomAndLinks[string(Links[i])], string(Links[j-1]))
				} else if j != len(Links)-1 {
					roomAndLinks[string(Links[i])] = append(roomAndLinks[string(Links[i])], string(Links[j+1]))
				}
			}
		}
		
	}

	// slFilter := []string{}

	// for _, l := range sl2 {
	// 	for _, l2 := range sl2 {
	// 	}
	// }

	fmt.Println(roomAndLinks)
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

	// fmt.Printf("%#v\n", lines)

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
			lines[i] = strings.ReplaceAll(lines[i], "-", "")
			if i != len(lines)-1 {
				Ants.Links += lines[i] + "*"
				// fmt.Println(len(lines)-1)
				// fmt.Println(i)
			} else {
				Ants.Links += lines[i]
			}

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
	fmt.Printf("%#v\n", Ants.Links)
	// fmt.Printf("%#v\n",Ants.Rooms)
	// fmt.Printf("%#v\n",Ants)
}
