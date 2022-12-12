package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our pizza b/w 1 and 10")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to ur rating: ", numRating+1)
	}

}
