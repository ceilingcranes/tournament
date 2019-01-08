package main

import (
	"fmt"
	"io/ioutil"
)

/* Describes each individual entry into the tournament. Each one will be compared to others in the tournament. */
type Contestant struct {
	ID   int
	Body string
}

/*
Each bracket is used to compare two contestants during the tournament. Each bracket contains two contestants, and contains the number of points for each.
*/
type Bracket struct {
	ID          int
	Contestant1 Contestant
	Contestant2 Contestant
	Points1     int
	Points2     int
}

/* Read in the data on each contestent from a file (where each one is seperated by a linebreak) and create brackets. The brackets are added to the list.
TODO: Using pointers vs just storing objects? */
func createBrackets(filename string) []Bracket {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))

	return []Bracket{}
}

func main() {
	var inputData = "testdata.txt"
	bracketList := createBrackets(inputData)
	fmt.Println(bracketList)
}
