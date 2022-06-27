package main

import "fmt"

func romanToInt(s string) int {

	var (
		lit           = make(map[rune]int)
		counter       = 0
		prevChar rune = ' '
		//number        = 0
		sum     = 0
		prevVal = 0
	)

	lit['I'] = 1
	lit['V'] = 5
	lit['X'] = 10
	lit['L'] = 50
	lit['C'] = 100
	lit['D'] = 500
	lit['M'] = 1000

	length := len(s)
	if length <= 0 || length > 15 {
		return 0 // Not sure what I need to return if error
	}

	for _, char := range s {
		if curVal, ok := lit[char]; ok {
			if prevVal, ok = lit[prevChar]; !ok { // First step
				prevChar = char
				sum += curVal
				continue
			}

			if curVal > prevVal {
				sum -= prevVal
				sum += curVal - prevVal
				counter = 0
			}
			if curVal < prevVal {
				sum += curVal
				counter = 0
			}
			if curVal == prevVal {
				switch char {
				case 'V', 'L', 'D':
					return 0
				default:
					if counter < 2 {
						sum += curVal
						counter++
					} else {
						return 0
					}
				}

			}
			prevChar = char

		} else {
			return 0 // Wrong literal
		}
	}
	return sum

}

func intToRoman(num int) string {
	var (
		lit           = make(map[int]rune)
		counter       = 0
		prevChar rune = ' '
		//number        = 0
		sum     = 0
		prevVal = 0
	)

	lit[1]    = 'I'
	lit[5]    = 'V'
	lit[10]   = 'X'
	lit[50]   = 'L'
	lit[1000] = 'C'
	lit[500]  = 'D'
	lit[1000] = 'M'

	if num > 0 && num < 3909 {



	}

}

func main() {

	var line string

	for {
		fmt.Println("Input Roman number:")
		fmt.Scanln(&line)
		fmt.Printf("Input %s\n", line)
		fmt.Printf("Return %d\n", romanToInt(line))
	}

}
