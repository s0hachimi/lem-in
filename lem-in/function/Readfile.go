package function

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile() Ant {
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
