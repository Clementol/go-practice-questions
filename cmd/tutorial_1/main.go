package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	fmt.Println(intNum + 1)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello World"
	fmt.Println(myString)

	fmt.Println((utf8.RuneCountInString(("YR"))))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intI int
	fmt.Println((intI))

	var1, var2 := 1, 2
	fmt.Println(var1, var2)
}
