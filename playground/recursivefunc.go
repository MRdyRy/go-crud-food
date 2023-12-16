// package main

// import "fmt"

// func factorialLoop(num int) int {
// 	result := 1

// 	for i := num; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

// //this is recursive function
// func factorialRecursive(num int) int {
// 	if num == 1 {
// 		return 1
// 	} else {
// 		return num * factorialRecursive(num-1)
// 	}
// }

// func main() {
// 	fmt.Println(factorialLoop(10))

// 	factor := factorialLoop

// 	fmt.Println(factor(10))

// 	fmt.Println(factorialRecursive(10))
// 	factor2 := factorialRecursive
// 	fmt.Println(factor2(10))
// }
