package main

import (
	"fmt"
	"github.com/DarkMagician55/tcolor"
)

func main() {
	fmt.Println(tcolor.Blue.Render("blue msg"))
	tcolor.NewTColor(tcolor.FgGreen, tcolor.BgBlue).Println("blue msg")
}
