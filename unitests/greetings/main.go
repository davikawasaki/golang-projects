package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	greetMessageEmpty := hello("")
	fmt.Println(aurora.Green(greetMessageEmpty))

	greetMessageJohn := hello("John")
	fmt.Println(aurora.Yellow(greetMessageJohn))
}
