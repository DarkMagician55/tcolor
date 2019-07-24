package main

import (
	"fmt"
	"github.com/DarkMagician55/tcolor"
)

func main() {
	//fmt.Println(tcolor.Red("red msg"))
	fmt.Println(tcolor.Render(tcolor.Red, "red msg"))
	fmt.Println(tcolor.Render(tcolor.Blue, "blue msg"))
	fmt.Println(tcolor.Blue.Render("blue msg"))
}
