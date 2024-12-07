package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func areNotTooFarApart(row []int) bool {
	for i := 0; i < (len(row) - 1); i++ {
		distance := IntAbs(row[i] - row[i+1])
		if distance < 1 || distance > 3 {
			return false
		}
	}

	return true
}

func convertToInt(line string) ([]int, error) {
	splitted := strings.Fields(line)
	result := make([]int, len(splitted))
	for i, elem := range splitted {
		num, err := strconv.Atoi(elem)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to int: %w", elem, err)
		}
		result[i] = num
	}

	return result, nil
}

func isStrictlyMonotonic(row []int, increasing bool) bool {
	for i := 0; i < len(row)-1; i++ {
		if increasing && row[i] >= row[i+1] {
			return false
		} else if !increasing && row[i] <= row[i+1] {
			return false
		}
	}
	return true
}

func isSafe(line string) (int, error) {
	convertedLine, err := convertToInt(line)
	if err != nil {
		return 0, err
	}
	if (isStrictlyMonotonic(convertedLine, true) || isStrictlyMonotonic(convertedLine, false)) && areNotTooFarApart(convertedLine) {
		return 1, nil
	}
	return 0, nil
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v\n", err)
	}

	result := 0
	for _, line := range lines {
		isSafeVal, err := isSafe(line)
		if err != nil {
			log.Printf("Error processing line '%s': %v\n", line, err)
			continue
		}
		result += isSafeVal
	}

	fmt.Printf("Number of safe lines: %d\n", result)
}
