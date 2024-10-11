package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	// *p = 1
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is %v\n", i)
	p = &i
	*p = 2
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is %v\n", i)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}