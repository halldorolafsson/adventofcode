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
		ps("Krav: https://adventofcode.com/2020/day/7")
	*/
	type bags struct {
		onebag         string
		containingBags string
	}

	//inits
	fileName := "input72.txt"
	splitOnContains := make([]string, 2)
	//	var rowCount int = 0
	//fileLength := lineCounter(fileName)
	//	allBags := make([]bags, fileLength)
	inputfile, err := os.Open(fileName)

	check(err)
	defer inputfile.Close()
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		//splitte på contain i to slice deler
		splitOnContains = strings.Split(scanner.Text(), "contain")
		fmt.Println(splitOnContains)
	}
	/*	for {
			work()
			if !condition {
				break
			}
		}
		for n := 0; n <= 5; n++ {
			if n%2 == 0 {
				continue
			}
			fmt.Println(n)
		}
	*/
}
func findItemsInSlice(mySlice []string, item string) int {
	var count int
	//containsShinyGoldBag := make([]string, 2)
	for _, sliceItem := range mySlice {
		if string(sliceItem) == item {
			ps(sliceItem)
			count++
		}
		//findItemsInSlice(mySlice, string(sliceItem[1]))
	}
	return count
}

func countInstance(str []rune) int {
	var count, result int
	for _, s := range str {
		count = strings.Count(string(str), string(s))
		if count > 1 {
			result = count
		}
	}
	return result
}
func checkKeyValue(key string, value string) bool {
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
func prSlice(s []rune) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
func lineCounter(filename string) int {
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
