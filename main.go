package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0.5),
		decimal.NewFromFloat(0.756),
		decimal.NewFromFloat(100),
		decimal.NewFromFloat(101),
		decimal.NewFromFloat(110),
		decimal.NewFromFloat(123),
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		response := convertNumbersToTH(input)
		fmt.Println(response)
	}
}

func convertNumbersToTH(input decimal.Decimal) string {
	// convert decimal to thai text (baht) and print the result here
	inputStr := input.String()
	//check for .
	splited := strings.Split(inputStr, ".")
	splitedNumbers := strings.Split(splited[0], "")

	response := ""
	locationRange := len(splitedNumbers)
	// fmt.Println(locationRange)
	for i, m := range splitedNumbers {
		response += assignNumberWord(m, locationRange-i, locationRange)
		response += assignLocationWord(m, locationRange-i)
	}

	if len(splited) == 2 {
		splitedDecimal := strings.Split(splited[1], "")
		decimalRange := 2
		response += "บาท"
		for i, m := range splitedDecimal {
			if i > 1 {
				continue
			}
			response += assignNumberWord(m, decimalRange-i, decimalRange)
			response += assignLocationWord(m, decimalRange-i)
		}
		response += "สตางค์"
	} else if len(splited) == 1 {
		response += "บาทถ้วน"
	} else {
		response = "error"
	}

	return response
}

func assignNumberWord(number string, location int, max int) string {
	switch number {
	case "0":
		if location == 1 && max == 1 {
			return "ศูนย์"
		} else {
			return ""
		}
	case "1":
		if location == 2 {
			return ""
		} else {
			return "หนึ่ง"
		}
	case "2":
		if location == 2 {
			return "ยี่"
		} else {
			return "สอง"
		}
	case "3":
		return "สาม"
	case "4":
		return "สี่"
	case "5":
		return "ห้า"
	case "6":
		return "หก"
	case "7":
		return "เจ็ด"
	case "8":
		return "แปด"
	case "9":
		return "เก้า"
	default:
		return ""
	}
}

func assignLocationWord(number string, location int) string {
	if number == "0" {
		return ""
	}
	switch location % 7 {
	case 1:
		return ""
	case 2:
		return "สิบ"
	case 3:
		return "ร้อย"
	case 4:
		return "พัน"
	case 5:
		return "หมื่น"
	case 6:
		return "แสน"
	default:
		return "ล้าน"
	}
}
