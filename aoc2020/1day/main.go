package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var linevalue int
var sum int = 2020

func main() {
	ps("Krav: Finne to verdier i datasettet som tilsammen en lik 2020")
	ps("Krav: Multiplisere tallene = svaret")
	ps("Åpne fil.")
	ps("Lese lineverdi (y) over i array")
	ps("Kjøre sjekk: 2020-y=z, der z er verdi i array i ny loop over tallsettet")
	ps("Multipliser tallene = = rett  svar")
	ps("Rydde opp hvis nødvendig")

	//linecounter -> primitive line count to create slice.
	//sliceOfValues := make([]int, linecounter("input1.txt"))
	sliceOfValues := make([]int, 0)

	//open the file
	inputfile, err := os.Open("input1.txt")
	check(err)

	//close the file later
	defer inputfile.Close()

	//create slice of values
	workscanner := bufio.NewScanner(inputfile)
	for workscanner.Scan() {
		linevalue, err := strconv.Atoi(workscanner.Text())
		check(err)
		//pi(linevalue)
		sliceOfValues = append(sliceOfValues, linevalue)
	}
	check(err)
	//pi(len(sliceOfValues))
	//pSlice(sliceOfValues)

	//find first result (produkt of the two numbers that sum to 2020)
	for j, v := range sliceOfValues {
		//ps(strconv.Itoa(j) + "  " + strconv.Itoa(v))
		for l, w := range sliceOfValues {
			//ps(strconv.Itoa(l) + " innerloop " + strconv.Itoa(w))
			if (v + w) == 2020 {
				ps(strconv.Itoa(v) + " + " + strconv.Itoa(w) + "=2020!")
				ps("Line: " + strconv.Itoa(j) + " and line: " + strconv.Itoa(l))
				r := v * w
				ps("Result 1: " + strconv.Itoa(r))
			}
		}
	}
	//find second result (produkt of the three numbers that sum to 2020)
	for k, x := range sliceOfValues {
		//ps(strconv.Itoa(j) + "  " + strconv.Itoa(v))
		for m, y := range sliceOfValues {
			for n, z := range sliceOfValues {
				if (x + y + z) == 2020 {
					ps(strconv.Itoa(x) + " + " + strconv.Itoa(y) + " + " + strconv.Itoa(z) + "=2020!")
					ps("Line: " + strconv.Itoa(k) + " and line: " + strconv.Itoa(m) + " and line: " + strconv.Itoa(n))
					s := x * y * z
					ps("Result 2: " + strconv.Itoa(s))
				}
			}
		}
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
