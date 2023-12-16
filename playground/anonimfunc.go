// package main

// import "fmt"

// //contract func
// type Blacklist func(string) bool

// // blacklist type Blacklist
// func registername(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("you are blocked ", name)
// 	} else {
// 		fmt.Println("welcome ", name)
// 	}
// }

// func main() {

// 	//anonim func inside var
// 	blacklistpeople := func(name string) bool {
// 		return name == "Anjing"
// 	}
// 	registername("john", blacklistpeople)

// 	registername("Anjing", blacklistpeople)

// 	//anonim function inside param
// 	registername("Doe", func(name string) bool {
// 		return "Anjing" == name
// 	})
// }
