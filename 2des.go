package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)


func main() {
plan("PLAN:")
plan("Åpne fil.")
plan("Lope over linjer og lesse linje for linjer i var  = line")
plan("Lage ny error handling funksjon")
plan("LInje line må sjekkes for to like og tre like bokstaver")
plan("loope igjennom line og hente ut en bokstav")
plan("loope igjennom line og hente ut en bokstav")


	file, err := os.Open("input_2.txt")
	check(err)
	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			fmt.Println(scanner.Text())
			check(err)
			line := "vtnihorkulbfvjcyzmsjgdxplw"

	}
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
