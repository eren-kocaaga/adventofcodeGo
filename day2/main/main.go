package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Sentence struct
{
	letter string
	minRep int
	maxRep int
	str string
}

func readData(path string) string{
	inputData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(inputData)
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func parseInput(inputStr string) []Sentence {
	inputStrArr := strings.Split(string(inputStr),"\n")

	var sentenceArr = []Sentence{}

	for _, i := range inputStrArr {
		sentenceStr := SplitAny(i,"- :")
		minRep_, err := strconv.Atoi(sentenceStr[0])
		maxRep_, err := strconv.Atoi(sentenceStr[1])
		if err != nil {
			panic(err)
		}
		sentenceTemp := Sentence{
			minRep: minRep_,
			maxRep: maxRep_,
			letter: sentenceStr[2],
			str: sentenceStr[3],
		}
		sentenceArr = append(sentenceArr, sentenceTemp)
	}
	return sentenceArr
}

func checkPassword(sentence Sentence) bool{
	count := strings.Count(sentence.str, sentence.letter)
	if count >= sentence.minRep && count <= sentence.minRep {
		return true
	}
	return false
}

func countValidPasswords(sentenceArr []Sentence) int{
	counter := 0
	for _, sentence := range sentenceArr{
		if checkPassword(sentence){
			counter++
		}
	}
	return counter
}

func main() {

	//retrieve Input
	inputStr := readData("/Users/erenkocaaga/go/src/adventOfCode/day2/main/input.txt")

	//parse Input
	sentenceArr := parseInput(inputStr)

	//count valid passwords
	start := time.Now()
	counter := countValidPasswords(sentenceArr)
	log.Printf("Count Valid Passwords took %s", time.Since(start))

	fmt.Println(counter)

}
