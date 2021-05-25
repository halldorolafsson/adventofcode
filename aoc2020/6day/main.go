package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	/*
		ps("Krav: https://adventofcode.com/2020/day/5")
	*/

	//inits
	fileName := "input6.txt"
	//fileLength := linecounter(fileName)
	var oneGroupAnswer string = ""
	alphabetmap := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
	sumOfUniques := 0

	inputfile, err := os.Open(fileName)
	check(err)

	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		check(err)
		if len(runes) > 0 {
			oneGroupAnswer = oneGroupAnswer + string(runes)
		}
		if len(runes) == 0 {
			uniques := 0
			for key, values := range alphabetmap {
				values = strings.Count(oneGroupAnswer, key)
				alphabetmap[key] = values
				if values != 0 {
					uniques++
				}
			}
			sumOfUniques = sumOfUniques + uniques
			oneGroupAnswer = ""
		}
	}
	fmt.Println("Part1=")
	fmt.Println(sumOfUniques)
}

func countInstance(str []rune) int {
	var l, unique, result int

	for _, s := range str {
		l = strings.Count(string(str), string(s))
		if l > 1 {
			result = l
		}
		if l == 1 {
			unique = unique + 1
		}
	}
	return result
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
