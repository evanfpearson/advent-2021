package solutions

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Day 1

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// Day 2

type Instruction struct {
	Direction string
	Distance int
}

func parse(instruction string) (Instruction, error) {
	s := strings.SplitN(instruction, " ", 2)
	dist, err := strconv.Atoi(s[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{
		Direction: s[0],
		Distance:  dist,
	}, nil
}

// Day 3

func Btoi(binary string) int {
	var i float64
	l := len(binary)
	for j := 1; j <= l; j++ {
		if string(binary[l-j]) == "1" {
			i += math.Pow(2, float64(j-1))
		}
	}
	return int(i)
}

func mostCommonBit(numbers []string, bitPosition int) (string, error) {
	var count [2]int
	for _, bits := range numbers {
		if string(bits[bitPosition]) == "1" {
			count[1]++
		} else if string(bits[bitPosition]) == "0" {
			count[0]++
		} else {
			return "", errors.New("not a binary string")
		}
	}
	if count[0] > count[1] {
		return "0", nil
	}
	return "1", nil
}

func leastCommonBit(numbers []string, bitPosition int) (string, error) {
	mcb, err := mostCommonBit(numbers, bitPosition)
	if err != nil {
		return "", err
	}
	if mcb == "1" {
		return "0", nil
	}
	return "1", nil
}

func getOxygenGeneratorRating(diagnosticReport []string) (int, error) {
	l := len(diagnosticReport)
	for bitPos := 0; l > 1; bitPos++ {
		var keep []string
		mcb, err := mostCommonBit(diagnosticReport, bitPos)
		if err != nil {
			return 0, err
		}
		for _, bits := range diagnosticReport {
			if string(bits[bitPos]) == mcb {
				keep = append(keep, bits)
			}
		}
		l = len(keep)
		diagnosticReport = keep
	}
	return Btoi(diagnosticReport[0]), nil
}

func getCO2ScrubberRating(diagnosticReport []string) (int, error) {
	l := len(diagnosticReport)
	for bitPos := 0; l > 1; bitPos++ {
		var keep []string
		lcb, err := leastCommonBit(diagnosticReport, bitPos)
		if err != nil {
			return 0, err
		}
		for _, bits := range diagnosticReport {
			if string(bits[bitPos]) == lcb {
				keep = append(keep, bits)
			}
		}
		l = len(keep)
		diagnosticReport = keep
	}
	return Btoi(diagnosticReport[0]), nil
}

