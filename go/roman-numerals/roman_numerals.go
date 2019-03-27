package romannumerals

import (
	"fmt"
)

// ToRomanNumeral coverts a normal number to a roman numeral
func ToRomanNumeral(num int) (romanNum string, err error) {
	if num <= 0 || num >= 3001 {
		return "", fmt.Errorf("Number %d is not valid, please enter a number between 1 and 3000", num)
	}

	for num > 0 {
		switch {
		case num >= 1000:
			romanNum += "M"
			num -= 1000
			continue
		case num >= 900:
			romanNum += "CM"
			num -= 900
			continue
		case num >= 500:
			romanNum += "D"
			num -= 500
			continue
		case num >= 400:
			romanNum += "CD"
			num -= 400
			continue
		case num >= 100:
			romanNum += "C"
			num -= 100
			continue
		case num >= 90:
			romanNum += "XC"
			num -= 90
			continue
		case num >= 50:
			romanNum += "L"
			num -= 50
			continue
		case num >= 40:
			romanNum += "XL"
			num -= 40
			continue
		case num >= 10:
			romanNum += "X"
			num -= 10
			continue
		case num >= 9:
			romanNum += "IX"
			num -= 9
			continue
		case num >= 5:
			romanNum += "V"
			num -= 5
			continue
		case num >= 4:
			romanNum += "IV"
			num -= 4
			continue
		case num >= 1:
			romanNum += "I"
			num--
			break
		}
	}

	return romanNum, nil
}
