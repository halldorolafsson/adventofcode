package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var sett, letter string
var min, max, i, count, valid1, mincount, maxcount, commoncount, valid2 int

func main() {
	/*
		ps("Krav: Finne to verdier i datasettet som tilsammen en lik 2020")
		ps("Krav: Multiplisere tallene = svaret")
		ps("Åpne fil.")
		ps("Lese lineverdi over i slice, parse line, konvetre datatyper, kjøre logikk")
		ps("Kjøre sjekk 1: finnes >=min og <=max bokstaver i ")
		ps("Kjøre sjekk 2: Er letter på rett plass? Og kun en gang? ")
		ps("Finne resultatet")
		ps("Rydde opp hvis nødvendig")
	*/
	//open the file
	inputfile, err := os.Open("input2.txt")
	check(err)

	//close the file later
	defer inputfile.Close()

	//timestamp
	fmt.Println(time.Now())

	//read file, create slice of values and extract elements
	workscanner := bufio.NewScanner(inputfile)
	for workscanner.Scan() {
		runes := []rune(workscanner.Text())
		//ps(string(runes))
		min, err := strconv.Atoi(strings.Trim(string(runes[0:strings.Index(string(runes), "-")]), " "))
		check(err)
		max, err := strconv.Atoi(strings.Trim(string(runes[strings.Index(string(runes), "-")+1:strings.Index(string(runes), "-")+3]), " "))
		check(err)
		letter := strings.Trim(string(runes[strings.Index(string(runes), ":")-1:strings.Index(string(runes), ":")]), " ")
		sett := strings.Trim(string(runes[strings.Index(string(runes), ":")+1:len(string(runes))]), " ")
		//ps(strconv.Itoa(min) + "-" + strconv.Itoa(max) + " " + letter + ": " + sett)
		check(err)
		count := 0
		//ps("***")
		//ps("min=" + strconv.Itoa(min))
		//ps("max=" + strconv.Itoa(max))
		//ps("letter=" + letter)
		//ps("sett=" + sett)

		//solution 1
		for i, s := range sett { //i start on 0!
			if letter == string(s) {
				count++
				//ps("i=" + strconv.Itoa(i))
				if (i + 1) == min {
					mincount++
					//ps("mincount=" + strconv.Itoa(mincount))
				}
				if (i + 1) == max {
					maxcount++
					//ps("maxcount=" + strconv.Itoa(maxcount))
				}
			}
		}
		if count >= min && count <= max {
			valid1++
		}
		//solution 2
		if letter == string(sett[min-1]) && letter == string(sett[max-1]) {
			commoncount++
		}
	}
	ps("valid1=")
	pi(valid1)
	valid2 = (mincount - commoncount) + (maxcount - commoncount)
	ps("valid2=")
	pi(valid2)
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
