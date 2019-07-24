package main

import (
	"fmt"
	"github.com/DarkMagician55/tcolor"
)

func main() {
	fmt.Println(tcolor.Blue.Render("blue msg"))
	fmt.Println(tcolor.NewTColor(tcolor.FgGreen, tcolor.BgBlue).Render("blue msg"))
}
