package main

import (
	"fmt"
	"strings"
)

func main() {
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("+-/*", r)
	}
	var inpStr string
	fmt.Scanf("%s\n", &inpStr)
	strValue := strings.FieldsFunc(inpStr, splitFunc)
	var numValue [2]int
	var isError bool
	var isRim bool
	operation := inpStr[len(strValue)-1]
	fmt.Printf("%c", operation)
	for idx, str := range strValue {
		isRim = true
		switch rv := str; rv {
		case "I":
			numValue[idx] = 1
			fmt.Println(numValue[idx], idx)
		case "II":
			numValue[idx] = 2
			fmt.Println(numValue[idx])
		default:
			isRim = false

			if idx > 0 {
				isError = true // если мы попали сюда не на первом элементе, то входящие данные не корректны
			}
		}
		if isRim == false {
			break
		}
	}
	print(isRim)
	print(isError)
	if isError == false && isRim == false {
		for idx, str := range strValue {

			switch av := str; av {
			case "1":
				numValue[idx] = 1
				fmt.Println(numValue[idx])
			case "2":
				numValue[idx] = 2
				fmt.Println(numValue[idx])
			default:
				isError = true // если мы попали сюда, то входящие данные не корректны
				fmt.Println(numValue[idx])

			}
			if isError == true {
				break
			}
		}
	}
	if isError == true {
		print("ОШИБОЧНЫЙ ВВОД!")
		return
	}

	switch operation {
	case '-':
		print(numValue[0] - numValue[1])
	case '+':
		print(numValue[0] + numValue[1])
	}

}
