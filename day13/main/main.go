package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func readData(path string) (string, error) {
	inputData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(inputData), nil
}

func parseInput(inputStr string) (int, []int, error) {
	in := strings.Split(string(inputStr), "\n")
	inn := strings.Split(string(in[1]), ",")
	buses := []int{}

	t, err := strconv.Atoi(in[0])
	if err != nil {
		return 0, buses, err
	}

	for _, val := range inn {
		if val != "x" {
			valI, err := strconv.Atoi(val)
			if err != nil {
				return t, buses, err
			}
			buses = append(buses, valI)
		}
	}
	return t, buses, nil
}

func findBus(dep int, buses []int) int {
	minTime := float64(dep)
	minBus := 0

	for _, val := range buses {
		res := float64(val) - math.Mod(float64(dep), float64(val))
		if res < minTime {
			minTime = res
			minBus = val
		}
	}

	return int(minTime) * minBus
}

func main() {

	start := time.Now()
	//defer log.Printf("total time %s", time.Now().Sub(start))

	inputStr, err := readData("./day13/main/input.txt")
	if err != nil {
		fmt.Printf("could not read data %s", err)
		return
	}

	t, b, err := parseInput(inputStr)
	if err != nil {
		fmt.Printf("could not parse data %s", err)
		return
	}

	res := findBus(t, b)
	log.Printf("total time %s", time.Now().Sub(start))
	fmt.Println(res)

}
