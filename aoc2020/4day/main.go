package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		ps("Krav: https://adventofcode.com/2020/day/4")
	*/

	var result1, result2, result3 int
	var onePassport string = ""
	var isonefalse bool

	//inits
	filename := "input4.txt"
	//fileLength := linecounter(filename)
	checkPassportFields := make([]string, 8)
	keyvalue := make([]string, 2)

	inputfile, err := os.Open(filename)
	check(err)
	defer inputfile.Close()

	//read row by row and do magic
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		check(err)
		if len(runes) > 0 {
			onePassport = onePassport + " " + string(runes)
		}
		if len(runes) == 0 {
			if (strings.Count(onePassport, ":") >= 7) && strings.Contains(onePassport, "byr") && strings.Contains(onePassport, "iyr") && strings.Contains(onePassport, "eyr") && strings.Contains(onePassport, "hgt") && strings.Contains(onePassport, "hcl") && strings.Contains(onePassport, "ecl") && strings.Contains(onePassport, "pid") {
				checkPassportFields = strings.Split(onePassport, " ")
				isonefalse = false
				result2 = 0
				for _, s := range checkPassportFields {
					if s != "" {
						keyvalue = strings.Split(s, ":")
						isonefalse = checkkeyvalue(keyvalue[0], keyvalue[1])
						if isonefalse == false && keyvalue[0] != "cid" {
							result2++
							break
						}
					}
				}
				if result2 == 0 { //Part two solution!
					result3++
				}
				//part one solution
				result1++
			}
			onePassport = ""
		}
	}
	ps("result1")
	pi(result1)
	ps("result3")
	pi(result3)
}

func checkkeyvalue(key string, value string) bool {
	var val int
	var result bool
	switch key {
	case "byr":
		val, _ = strconv.Atoi(value)
		if val >= 1920 && val <= 2002 {
			result = true
		}
	case "iyr":
		val, _ = strconv.Atoi(value)
		if val >= 2010 && val <= 2020 {
			result = true
		}
	case "eyr":
		val, _ = strconv.Atoi(value)
		if val >= 2020 && val <= 2030 {
			result = true
		}
	case "hgt":
		heightandsystem := make([]string, 2)
		//metric cm
		if strings.Contains(value, "cm") {
			heightandsystem = strings.Split(value, "cm")
			height, err := strconv.Atoi(heightandsystem[0])
			check(err)
			if height >= 150 && height <= 193 {
				result = true
			}
			//inches
		}
		if strings.Contains(value, "in") {
			heightandsystem = strings.Split(value, "in")
			height, err := strconv.Atoi(heightandsystem[0])
			check(err)
			if height >= 59 && height <= 76 {
				result = true
			}
		}
	case "hcl":
		if strings.HasPrefix(value, "#") {
			if len(value) == 7 {
				for _, s := range value {
					if strings.Compare(string(s), "g") == -1 {
						result = true
					} else {
						result = false
					}
				}
			}
		}
	case "ecl":
		if strings.Contains(value, "amb") || strings.Contains(value, "blu") || strings.Contains(value, "brn") || strings.Contains(value, "gry") || strings.Contains(value, "grn") || strings.Contains(value, "hzl") || strings.Contains(value, "oth") {
			result = true
		}
	case "pid":
		if len(value) == 9 {
			for _, s := range value {
				if strings.Compare(string(s), "a") == -1 {
					result = true
				} else {
					result = false
				}
			}
		}
	case "cid":
		//result = true
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
