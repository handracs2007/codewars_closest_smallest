// https://www.codewars.com/kata/5868b2de442e3fb2bb000119/train/go

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Number struct {
	DigitStr string
	Weight   int
	Position int
}

func Weight(number string) int {
	var num, _ = strconv.Atoi(number)
	var weight = 0

	for num > 0 {
		weight += num % 10
		num /= 10
	}

	return weight
}

func Closest(strng string) string {
	if len(strng) == 0 {
		// handle empty string
		return "[(), ()]"
	}

	// first, let's get the list of numbers by splitting the string by space
	var splitString = strings.Split(strng, " ")
	var numbers = make([]Number, len(splitString))

	for idx, val := range splitString {
		var number = Number{
			DigitStr: val,
			Weight:   Weight(val),
			Position: idx,
		}

		numbers[idx] = number
	}

	// second, let's sort the numbers by weight and position
	sort.Slice(numbers, func(i, j int) bool {
		var number1 = numbers[i]
		var number2 = numbers[j]

		// if weight is different, sort by weight
		if number1.Weight != number2.Weight {
			return number1.Weight < number2.Weight
		}

		// if weight is the same, sort by position
		return number1.Position < number2.Position
	})

	// third, loop through each element and find the position of 2 items with smallest difference of weight
	var minPos1, minPos2, minDiff = -1, -1, math.MaxInt64

	for i := 1; i < len(numbers); i++ {
		var diff = numbers[i].Weight - numbers[i-1].Weight

		if diff < minDiff {
			minDiff = diff
			minPos1 = i - 1
			minPos2 = i
		}
	}

	// fourth, build the result
	var number1 = numbers[minPos1]
	var number2 = numbers[minPos2]
	return fmt.Sprintf("[(%d, %d, %s), (%d, %d, %s)]", number1.Weight, number1.Position, number1.DigitStr, number2.Weight, number2.Position, number2.DigitStr)
}

func main() {
	fmt.Println(Closest(""))
	fmt.Println(Closest("456899 50 11992 176 272293 163 389128 96 290193 85 52"))
	fmt.Println(Closest("239382 162 254765 182 485944 134 468751 62 49780 108 54"))
}
