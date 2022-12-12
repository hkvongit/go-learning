package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my app"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a rating to our pizza:")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating, rating: ", input)
	fmt.Printf("Data type of rating: %T", input)
}
