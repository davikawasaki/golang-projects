package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	greetMessage := hello("John")
	fmt.Println(aurora.Yellow(greetMessage))
}
