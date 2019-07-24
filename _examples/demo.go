package main

import (
	"fmt"
	"github.com/DarkMagician55/tcolor"
)

func main() {
	fmt.Println(tcolor.Red("red msg"))
	fmt.Println(tcolor.Render(tcolor.FgBlue, "red msg"))
}
