package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Naml struct {
	Nmala int
	Start string
	End   string
	Rooms []string
	Links []string
}

var Graph struct {
	Path []string
}


func main() {
	Readfile()
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
	Naml.Nmala = n

	fmt.Printf("%#v\n", lines)

	for i := 0; i <= len(lines)-1; i++ {

		if len(lines[i]) != 0 && string(lines[i][0]) == "#" && string(lines[i][1]) != "#" {
			continue
		}

		if lines[i] == "##start" {
			for j := 0; j <= len(lines[i+1])-1; j++ {
				if lines[i+1][j] == ' ' {
					Naml.Start = string(lines[i+1][:j])
                    break
				}
			}
		}

		if lines[i] == "##end" {
			for j := 0; j <= len(lines[i+1])-1; j++ {
				if lines[i+1][j] == ' ' {
					Naml.End = string(lines[i+1][:j])
                    break
				}
			}
		}

		if strings.Contains(lines[i], "-") {
			Naml.Links = append(Naml.Links, lines[i])
		} else {
			for j := 0; j <= len(lines[i])-1; j++ {
				if lines[i][j] == ' ' {
					Naml.Rooms = append(Naml.Rooms,  string(lines[i][:j]))
                    break
				}
			}
		}

	}
	// fmt.Println(Naml.Nmala)
	// fmt.Println(Naml.Start)
	// fmt.Println(Naml.End)
    // fmt.Printf("%#v\n",Naml.Links)
	// fmt.Printf("%#v\n",Naml.Rooms)
	// fmt.Printf("%#v\n",Naml)
}
