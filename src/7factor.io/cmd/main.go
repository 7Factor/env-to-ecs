package main

import "fmt"

func main() {
	fmt.Println("The main func")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}