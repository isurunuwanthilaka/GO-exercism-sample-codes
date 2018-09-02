package raindrops

import (
	"strconv"
)

func Convert(num int) string {
	result := ""
	check := 1
	if num%3 == 0 {
		result += "Pling"
		check = 0
	}
	if num%5 == 0 {
		result += "Plang"
		check = 0
	}
	if num%7 == 0 {
		result += "Plong"
		check = 0
	}
	if check == 1 {
		result = strconv.Itoa(num)
	}
	return result
}
