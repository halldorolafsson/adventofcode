package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	/*
		ps("Krav: https://adventofcode.com/2020/day/5")
	*/

	//inits
	filename := "input5.txt"
	//fileLength := linecounter(filename)
	//frontback := make([]string, fileLength)
	//leftright := make([]string, fileLength)
	//keyvalue := make([]string, 2)
	var seatRow = make([]int, 128)
	var initRow = make([]int, 128)
	var seatCol = make([]int, 8)
	var initCol = make([]int, 8)
	var row, column, seatID, highestSeatID int

	for l := 0; l < 128; l++ {
		initRow[l] = l
	}
	for l := 0; l < 8; l++ {
		initCol[l] = l
	}

	inputfile, err := os.Open(filename)
	check(err)

	defer inputfile.Close()

	//read row by row and do magic
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		check(err)
		//ps(string(runes))
		seatRow = initRow
		seatCol = initCol
		for _, s := range runes {
			//ps(string(s))
			if string(s) == "F" || string(s) == "B" {
				if string(s) == "F" {
					seatRow = seatRow[:len(seatRow)/2]
				}
				if string(s) == "B" {
					seatRow = seatRow[(len(seatRow) / 2):]
				}
			}
			if len(seatRow) == 1 {
				row = seatRow[0]
			}
			if string(s) == "R" || string(s) == "L" {
				if string(s) == "L" {
					seatCol = seatCol[:len(seatCol)/2]
				}
				if string(s) == "R" {
					seatCol = seatCol[(len(seatCol) / 2):]
				}
			}
			if len(seatCol) == 1 {
				column = seatCol[0]
			}
		}
		//		pSlice(seatRow)
		seatID = (row * 8) + column
		fmt.Println("seatID = " + strconv.Itoa(seatID))
		if seatID > highestSeatID {
			highestSeatID = seatID
			//fmt.Println("highestSeatID = " + strconv.Itoa(highestSeatID))
		}
	}
}

func checkkeyvalue(key string, value string) bool {
	var result bool
	switch key {
	case "1":
	case "2":
	case "3":
	}
	return result
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
func psSlice(s []string) {
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
