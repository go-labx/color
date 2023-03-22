package main

import (
	"fmt"
	"github.com/go-labx/color"
)

func main() {
	fmt.Println(color.BlackString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.RedString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.GreenString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.YellowString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.BlueString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.MagentaString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.CyanString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
	fmt.Println(color.WhiteString("username = %s age = %d city = %s", "Jack", 20, "Hangzhou"))
}
