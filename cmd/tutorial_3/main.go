package main

import "fmt"

// Arrays, Slices, Maps and Loops
func main() {
	var intArr [3]int32 = [...]int32{1, 2, 3}
	// intArr := [...]int32{1,2,3}
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	intSlice = append(intSlice, 8)
	intSlice = append(intSlice, 9)
	intSlice = append(intSlice, 10)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	fmt.Println(intSlice)
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	var intSlice3 []int32 = make([]int32, 3)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]
	// delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
	// Go doesn't have while loop
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
