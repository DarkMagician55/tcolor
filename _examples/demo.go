package main

import (
	"fmt"
	"github.com/DarkMagician55/tcolor"
	"log"
)

func main() {
	//fmt.Println(tcolor.Red("red msg"))
	fmt.Println(tcolor.Render(tcolor.Red, "red msg"))
	log.Println(tcolor.Render(tcolor.Blue, "blue msg"))
}
