package main

import (
	"fmt"
	"os"
)

var symbol = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func getSymbolValue(val uint8) int {
	return symbol[fmt.Sprintf("%c", val)]
}

func romanToInt(str string) int {

	//строка как массив берем ее длину а если кириилица то больше
	lastIndex := len(str) - 1
	value := getSymbolValue(str[lastIndex])

	//fmt.Println(lastIndex)
	//fmt.Println(value)

	for i := 0; i < lastIndex; i++ {
		int1 := getSymbolValue(str[i])
		int2 := getSymbolValue(str[i+1])
		fmt.Println(int1)
		fmt.Println(int2)

		if int1 >= int2 {
			value += int1
		} else {
			if int1*5 == int2 || int1*10 == int2 {
				value -= int1
			}
		}
	}

	return value
}

func main() {

	arguments := os.Args

	fmt.Println(arguments[1])
	//fmt.Println(arguments)

	//strs := strings.Split(arguments[1], ",")

	res := romanToInt(arguments[1])
	fmt.Println(res)
}
