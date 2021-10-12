package main

import "fmt"

func main() {
	//var whattosay string (one way for Declaration)
	//whattosay = "Hello world"
	whattosay := "Hello world" //another way of Declaration 
	new(whattosay)
} 
func new(whattosay string) {
	fmt.Println(whattosay)

}