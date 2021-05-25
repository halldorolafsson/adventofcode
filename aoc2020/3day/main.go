package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var row, column, right, down, trees int
var tree string

func main() {
	/*
		ps("Krav:")
	*/

	//inits
	trees := 0
	row := 0
	column := 0
	right := 3
	down := 1
	tree := '#'
	filename := "input3.txt"

	//always usefull
	filelength := linecounter(filename)

	//open the file for processing
	inputfile, err := os.Open(filename)
	check(err)

	//close the file later
	defer inputfile.Close()

	//timestamp
	fmt.Println(time.Now())

	//read file, create slice of values and extract elements
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		check(err)
		if column >= len(runes) {
			column = 0
		}
		if row >= filelength {
			break
		}
		if runes[column] == tree {
			trees++
		}
		column = column + right
		row = row + down
	}
	ps("trees = " + strconv.Itoa(trees))
}
func ps(phrase string) {
	fmt.Println(phrase)
}
func pi(number int) {
	fmt.Println(number)
}
func pSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
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
