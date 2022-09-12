package main

import "fmt"

type Person struct {
	Name string
	Sex string
	Age int
	Height int
	Weight int
}

func main() {
	fmt.Println(Person{"Musashi", "男", 27, 182, 74})
}
