// package main

// import "fmt"

// // function sebagai parameter
// func filterword(name string, filter func(string) string) {
// 	filteredword := filter(name)
// 	fmt.Println(filteredword)
// }

// //function same as contract
// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	//call func
// 	filterword("Anjing", spamFilter)

// 	//call with param
// 	varparam := spamFilter
// 	filterword("Anjir", varparam)
// }
