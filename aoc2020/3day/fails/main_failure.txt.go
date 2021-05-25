package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var space, tree string
var i, count, linecount, constant int

func main() {
	/*
		ps("Krav:Naviger 3 til høyre og en ned gjennom hele inputfila. Punktum (.) er åpent rom men skigrad (#) er et tre")
		ps("Krav: Tell antall tre")
		ps("Åpne fil.")
		ps("Lese linje over i slice lese tegng for tegn og telle")
	*/
	//open the file
	inputfile, err := os.Open("input32.txt")
	check(err)

	//close the file later
	defer inputfile.Close()

	//timestamp
	fmt.Println(time.Now())

	constant := 34

	//read file, create slice of values and extract elements
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		//ps(string(runes))
		check(err)
		//ps("linecount=" + strconv.Itoa(linecount))
		for i, s := range runes {
			//ps("i=" + strconv.Itoa(i))
			count++
			if i >= 30 {
				linecount++
				ps("linecount=" + strconv.Itoa(linecount))
			}
			if linecount >= 1 {
				ps("count=" + strconv.Itoa(count))
				if count == constant*linecount {
					ps("s=" + string(s))
				}
			}
		}
	}
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
