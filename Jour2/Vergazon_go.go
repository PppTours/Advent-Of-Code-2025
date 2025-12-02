// To run this program :
// From the Jour2 directory (input data path is hardcoded)
// > go run Vergazon_go.go

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func partie1(ranges []Range) {

	// Next step : detect invalid IDs and calculate their sum !
	//
	// Since the young Elf was just doing silly patterns, you can find the invalid IDs by looking
	// for any ID which is made only of some sequence of digits repeated twice.
	// So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.
	//
	// None of the numbers have leading zeroes; 0101 isn't an ID at all. (101 is a valid ID that you would ignore.)
	//
	// This means the naive method would be to iterate trough all the IDs, trim the zeros, and
	// check for validity here.
	// Would there be a faster way to do it ?
	// First of all, an invalid ID always has an even number of digits.
	// Maybe a smarter way to do this would be to generate a plausible list of invalid IDs, and
	// Check if they actually belong in the range ?
	//
	// For example : let's say we have the range 95-115
	// Let's take the nearest "double" digit below 95 -> 88
	// Let's take the nearest "double" digit above 115 -> 1010
	// We now have a list of "doubled" digits to check : 8, 9, 10.
	// We can generate their "doubled" value -> 88, 99, 1010 and check if they belong in the range.
	// This avoids checking every single value in the list.
	// However, taking the nearest "double" digit above or below a number might be longer
	// than iterating trough every number in the range.
	//
	//
	// Let's stick to the naive version for now.

	var password int = 0

	for _, r := range ranges { // For each range
		for i := r.Start; i <= r.End; i++ { // For each index in the range
			// Convert the number to a string
			// Warning : don't use string(i), it will treat i as an ASCII code and return a single char.
			var i_string string = strconv.Itoa(i)
			// First check : is the number of characters even ?
			var is_even bool = (len(i_string) % 2) == 0
			if !is_even {
				continue
			}
			// Second check : is the first half equal to the second half ?
			n := len(i_string) // length of the string
			mid := n / 2       // midpoint
			firstHalf := i_string[:mid]
			secondHalf := i_string[mid:]
			if firstHalf == secondHalf {
				password += i
			}
		}
	}

	fmt.Printf("PARTIE 1 | Password : %d\n", password)
	// 15873079081
}

func partie2(ranges []Range) {
	fmt.Printf("PARTIE 2")

	// Now, an ID is invalid if it is made only of some sequence of digits repeated
	// at least twice.
	// So, 12341234 (1234 two times),
	// 	123123123 (123 three times),
	// 	1212121212 (12 five times),
	// 	and 1111111 (1 seven times)
	// are all invalid IDs.

	// This complicates the problem a lot.
	// We need to get all the configurations a number could be repeated.
	// For example, 8 can be 8 times 1, 4 times 2 or 2 times 4.
	// Basically, we need all the number that divide the character length.

	var password int = 0
	for _, r := range ranges { // For each range
		for i := r.Start; i <= r.End; i++ { // For each index in the range
			// Convert the number to a string
			var i_string string = strconv.Itoa(i)
			var n int = len(i_string)
			var divisors []int

			// Get divisors
			// Iterate from 1 up to n
			for i := 1; i <= n; i++ {
				// If n is divisible by i (remainder is 0)
				if n%i == 0 {
					divisors = append(divisors, i)
				}
			}

			is_valid := true
			for _, divisor := range divisors {
				// 1. FIX: Skip the case where the divisor is 1
				// (A string is always equal to itself, so we ignore this case)
				if divisor == 1 {
					continue
				}

				// Reconstruct the string with the repeated pattern and check if it matches
				partLength := n / divisor
				pattern := i_string[0:partLength]
				reconstructed := strings.Repeat(pattern, divisor)

				// If it matches, the ID is a repetition.
				if reconstructed == i_string {
					is_valid = false
					break
				}
			}

			if !is_valid {
				password += i
			}

		}
	}

	fmt.Printf("PARTIE 2 | Password : %d\n", password)
	// 22617871034

}

func main() {
	// First step : read the file and load data
	content, err := os.ReadFile("Vergazon_inputData.txt")
	if err != nil {
		fmt.Println("File could not be read:", err)
		return
	}
	// Remove whitespaces
	var fileContent string = strings.TrimSpace(string(content))
	var ranges []Range
	// Split by commas
	var rawRanges []string = strings.Split(fileContent, ",")
	for _, rr := range rawRanges {
		// Split each range by the "-" character
		parts := strings.Split(rr, "-")

		if len(parts) != 2 {
			fmt.Printf("Skipping invalid format: %s\n", rr)
			continue
		}
		// Convert strings to integers
		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Printf("Error converting numbers in range: %s\n", rr)
			continue
		}
		ranges = append(ranges, Range{Start: start, End: end})
	}

	// We now have a nice array of Range structs.
	fmt.Printf("Successfully parsed %d ranges:\n", len(ranges))
	for _, r := range ranges {
		fmt.Printf("Start: %d, End: %d\n", r.Start, r.End)
	}

	partie1(ranges)
	partie2(ranges)

}
