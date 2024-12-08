package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v\n", err)
	}

	result := 0

	for _, line := range lines {
		occurences := pattern.FindAllString(line, -1)
		for _, elem := range occurences {
			elem = strings.Replace(elem, "mul(", "", -1)
			elem = strings.Replace(elem, ")", "", -1)
			numbers := strings.Split(elem, ",")
			x, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatalf("Error converting input to integer: %v\n", err)
			}
			y, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Fatalf("Error converting input to integer: %v\n", err)
			}
			result += x * y
		}
	}

	fmt.Println(result)
}
