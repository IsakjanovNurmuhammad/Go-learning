package main

import (
	"fmt"
	"unicode/utf8"
)

// Q2.1
//
//	func main() {
//		for i := 10; i < 10; i++ {
//			fmt.Printf("%d\n", i)
//		}
//	}
//

// Q2.2
//
//	func main() {
//		num := 0
//
// here:
//
//		fmt.Printf("%d\n", num)
//		if num < 10 {
//			num++
//			goto here
//		}
//	}

// Q2.3
//func main() {
//	var arr [10]int
//	for num := 0; num < 10; num++ {
//		arr[num] = num
//	}
//	fmt.Printf("%d", arr)
//}

// Q3
//func main() {
//	const (
//		FIZZ = 3
//		BUZZ = 5
//	)
//	for i := 1; i < 100; i++ {
//		if i%FIZZ == 0 {
//			fmt.Print("Fizz")
//		}
//		if i%BUZZ == 0 {
//			fmt.Print("Buzz")
//		}
//		fmt.Println()
//	}
//}

// Q4.1
//func main() {
//	str := "A"
//	for i := 0; i < 100; i++ {
//		str = str + "A"
//		fmt.Println(str)
//	}
//}

// Q4.2
func main() {
	str := "A"
	length := 0
	for i := 0; i < 100; i++ {
		str = str + "A"
		length = length + i
		fmt.Println(str)
	}
	fmt.Printf("%d\n", length)
	fmt.Printf("Runes: %d\n",
		len([]byte(str)), utf8.RuneCount([]byte(str)))
}

//Q4.3
