package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	s "strings"
)


func main() {
plan("PLAN:")
plan("Åpne fil.")
plan("Lope over linjer og lesse linje for linjer i var  = line")
plan("Lage ny error handling funksjon")
plan("LInje line må sjekkes for to like og tre like bokstaver")
plan("loope igjennom line og hente ut en og en bokstav")
plan("telle antall like bokstaver i linja")
plan("først to like så tre like")
plan("lage sjekksum. Dele toer tallet på to og treertallet på tre")
plan("ferdig")
plan("brukt side: https://gobyexample.com og https://tour.golang.org")

	file, err := os.Open("input_2.txt")
	check(err)
	defer file.Close()
	var twos, threes int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			runes := []rune(scanner.Text())
			line := string(runes[0:len(scanner.Text())])
			//fmt.Println(line)
			check(err)
			for i := 0; i < len(line);i++ {
					letter := line[i:i+1]
					check(err)
					if s.Count(line,letter) == 2 {
						twos++
					}
					if s.Count(line,letter) == 3 {
						threes++
					}
					//if s.Count(line,letter) >= 2 {
						//fmt.Println(s.Count(line,letter), letter,twos/2,threes/3)
					//}
			}
	}
	checksum := ((twos/2) * (threes/3))
	fmt.Println(checksum)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
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
