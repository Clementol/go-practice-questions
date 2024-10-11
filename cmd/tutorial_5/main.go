package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[2]
	fmt.Printf("%v, %T\n", indexed, indexed)
	fmt.Println(myString)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Println(len(myString))

	t0 := time.Now()
	var strSlice = []string{"S", "U", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	// var catStr string = ""
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
		// catStr += strSlice[i]
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)
	fmt.Println(time.Since(t0))
}
