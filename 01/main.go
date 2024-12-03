package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func separateLocationIds(locationIds []string) ([]int, []int) {
	left := make([]int, len(locationIds))
	right := make([]int, len(locationIds))

	for _, line := range locationIds {
		splitted := strings.Split(line, "   ")
		first, err := strconv.Atoi(splitted[0])
		if err != nil {
			fmt.Printf("error converting the left locationId to a string. Error: %v", err)
			return nil, nil
		}
		second, err := strconv.Atoi(splitted[1])
		if err != nil {
			fmt.Printf("error converting the right locationId to a string. Error: %v", err)
			return nil, nil
		}
		left = append(left, first)
		right = append(right, second)
	}

	return left, right
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func computeDistance(left, right []int) int {
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance = totalDistance + IntAbs(left[i]-right[i])
	}

	return totalDistance
}

func countOccurances(arr []int) map[int]int {
	freq := make(map[int]int)

	for _, num := range arr {
		freq[num]++
	}

	return freq
}

func computeSimilarity(arr []int, m map[int]int) int {
	similarity := 0

	for _, elem := range arr {
		similarity = similarity + elem*m[elem]
	}

	return similarity

}

func main() {
	lines := readInput()
	left, right := separateLocationIds(lines)
	sort.Ints(left)
	sort.Ints(right)
	distance := computeDistance(left, right)
	fmt.Println("Part 1:")
	fmt.Printf("Total distance: %d\n", distance)
	fmt.Println("-----------------------------------")
	fmt.Println("Part 2:")
	occurances := countOccurances(right)
	similarity := computeSimilarity(left, occurances)
	fmt.Printf("Similarity score: %d\n", similarity)
}
