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
	const LEN = 13
	var (
		result string
		lit    = make(map[int]string)
		nums   = [LEN]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	)

	lit[1] = "I"
	lit[4] = "IV"
	lit[5] = "V"
	lit[9] = "IX"
	lit[10] = "X"
	lit[40] = "XL"
	lit[50] = "L"
	lit[90] = "XC"
	lit[100] = "C"
	lit[400] = "CD"
	lit[500] = "D"
	lit[900] = "CM"
	lit[1000] = "M"

	for j := 0; j < LEN; j++ {
		if num > 0 && num < 4000 {
			res := num / nums[j]
			if res > 3 {
				result = result + lit[nums[j]]
				result = result + lit[nums[j-1]]
			} else {
				for i:=res;i>0;i-- {
					result = result + lit[nums[j]]
				}
			}
			num-=nums[j]*res
		} else { break }
	}
	return result
}

func main() {

	var line string

	for {
		fmt.Println("Input Roman number:")
		fmt.Scanln(&line)
		fmt.Printf("Input %s\n", line)
		r := romanToInt(line)
		fmt.Printf("Integer Return %d\n", r)
		fmt.Printf("Back to Roman: %s\n", intToRoman(r))
	}

}
