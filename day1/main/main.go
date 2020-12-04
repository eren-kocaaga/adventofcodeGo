package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums{
		if val, ok := m[target-n]; ok {
			return []int{val,i}
		}
		m[n]=i
	}
	return nil
}

func readData(path string) string{
	inputData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(inputData))

	return string(inputData)
}


func parseInput(inputStr string) []int {
	inputStrArr := strings.Split(inputStr,"\n")

	var inputIntArr = []int{}

	for _, i := range inputStrArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputIntArr = append(inputIntArr, j)
	}

	return inputIntArr
}

func main() {

	//retrieve Input
	inputStr := readData("/Users/erenkocaaga/go/src/adventOfCode/day1/main/input.txt")

	//parse Input
	dataIntArr := parseInput(inputStr)

	//process Data
	start := time.Now()
	res := twoSum(dataIntArr, 2020)
	log.Printf("twoSum took %s", time.Since(start))

	//output Result
	fmt.Println(res[0], "-", res[1])
	fmt.Println(dataIntArr[res[0]], "-", dataIntArr[res[1]])



}



