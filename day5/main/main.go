package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func readData(path string) string{
	inputData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(inputData)
}

func compare(first string, second string) string {
    if first == "" { return second }

    for index, char := range first {
        if char == 'R' && second[index] != 'R' {
            return first
        } else if char != 'R' && second[index] == 'R' {
            return second
        } else if char == 'B' && second[index] != 'B' {
            return first
        } else if char != 'B' && second[index] == 'B' {
            return second
        }
    }

    return first
}

func parseInput(inputStr string) []string {
	inputStrArr := strings.Split(string(inputStr),"\n")
    return inputStrArr
}


func getId(in string) int64 {
    rv := int64(0)

    for _, char := range in {
        rv = rv << 1

        switch char {
            case 'R': rv = rv + 1
            case 'B': rv = rv + 1
            default:
        }
    }

    return rv
}

func getHighest(arr []string) int64 {

    highest := arr[0]
    for _, data := range arr {
        highest = compare(data, highest)
    }
    return getId(highest)
}

func getHighestSlow(arr []string) int64 {
    max := int64(0)
    for _, data := range arr {
        if id := getId(data); id > max {
            max = id
        }
    }
    return max
}


func main() {
	start := time.Now()

	inputStr := readData("./day5/main/input.txt")
	arr := parseInput(inputStr)

	fmt.Println(getHighest(arr))
	end := time.Now()

	log.Printf("total time %s", end.Sub(start))

}
