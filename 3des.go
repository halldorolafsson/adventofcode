package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
  "strconv"
	s "strings"

)
// to store claim data
type claim struct {
	id int
	leftstart int
	topstart  int
	leftwide int
	toptall int
}

func main() {
	plan("PLAN:")
	//linecounter -> primitive line count to create slice/array. Used to define slice that will store the claims
	sliceOfclaims := make([]claim, linecounter("input_3.txt"))

	//open the file
	input_file, err := os.Open("input_3.txt")
	check(err)

	//close the file later
	defer input_file.Close()

	//start working, read file line by line
	workscanner := bufio.NewScanner(input_file)
	for workscanner.Scan() {
		//rune is a slice (read more)
		runes := []rune(workscanner.Text())
		//convert slice to string -> line
		line := string(runes[0:len(workscanner.Text())])
		//parse each line and find what we are looking for
		sid, err :=   strconv.Atoi(line[1:s.Index(line,"@")-1])
		sleftstart,err := strconv.Atoi(line[s.Index(line,"@")+2:s.Index(line,",")])
		stopstart, err := strconv.Atoi(line[s.Index(line,",")+1:s.Index(line,":")])
		sleftwide, err := strconv.Atoi(line[s.Index(line,":")+2:s.Index(line,"x")])
		stoptall, err := strconv.Atoi(line[s.Index(line,"x")+1: len(line)])
		check(err)

		// add struct claim to sliceOfclaims
		sliceOfclaims = append(sliceOfclaims, claim{sid, sleftstart, stopstart, sleftwide, stoptall})
		// investigate if this should bed stored in a map
		// create a map and add sliceOfclaims

	}
	fmt.Println(sliceOfclaims)

	if err := workscanner.Err(); err != nil {
			log.Fatal(err)
	}
}


func linecounter(filename string) int {
	file, err := os.Open(filename)
	check(err)
	l := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l++
	}
	file.Close()
	return l
}


func plan(phrase string) {
	fmt.Println(phrase)
}

func check(e error) {
    if e != nil {
			log.Fatal(e)
      panic(e)
    }
}
