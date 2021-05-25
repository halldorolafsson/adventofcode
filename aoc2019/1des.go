package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
plan("PLAN:")
plan("Åpne fil.")
plan("Lope over linjer og lese verdier")
plan("Konvertere string verdier til integer ")
plan("Finne resultat")
plan("Lukke fil")

	file, err := os.Open("input.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var y int
	for scanner.Scan() {
			//fmt.Println(scanner.Text())
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
					log.Fatal(err)
			}
			fmt.Println(i)
			y +=i
			fmt.Println("Y er lik= ", y)
}
	fmt.Println(y)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func plan(phrase string) {
	fmt.Println(phrase)
}
