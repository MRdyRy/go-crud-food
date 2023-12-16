// package main

// import (
// 	"fmt"
// )

// // func name param = string, varags, dengan multiple return yaitu string & int
// func getName(nama string, numbers ...int) (string, int) {

// 	total := 0
// 	//menghiraukan value menggunakan keyword _
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	//multiple return
// 	return nama, total
// }

// func main() {
// 	//variable namanya menyimpan value sebuah function
// 	namanya := getName

// 	//memanggil variable yang sudah menjadi function
// 	fmt.Println(namanya("jondoe", 1, 1, 1, 1, 1))
// }
