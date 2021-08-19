package main

import "fmt"

func main() {
	baseArray := [...]int{1, 2, 3} // array of length 3
	fmt.Printf("%#v\n\n", baseArray)

	slice1 := baseArray[:]
	fmt.Printf("%#v\n", slice1)
	fmt.Printf("slice1 len: %d\ncap: %d\n\n", len(slice1), cap(slice1))

	slice2 := append(slice1, 4, 5)
	fmt.Printf("slice1 = %#v\nslice2 = %#v\n\n", slice1, slice2)

	baseArray[0] = 100
	fmt.Printf("baseArray = %#v\nslice1 = %#v\nslice2 = %#v\n\n", baseArray, slice1, slice2)

	slice3 := make([]int, 3, 4)
	fmt.Printf("slice3 = %#v\nlen = %d\ncap = %d\n\n", slice3, len(slice3), cap(slice3))

	slice4 := append(slice3, 1)
	fmt.Printf("slice3 = %#v\nlen = %d\ncap = %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4 = %#v\nlen = %d\ncap = %d\n\n", slice4, len(slice4), cap(slice4))

	slice5 := append(slice3, 2)
	fmt.Printf("slice3 = %#v\n", slice3)
	fmt.Printf("slice4 = %#v !!!\n", slice4)
	fmt.Printf("slice5 = %#v\n\n", slice5)
}
