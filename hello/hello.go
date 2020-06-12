package hello

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// SayHelloAndOneMoreThing function: Says hello
func SayHelloAndOneMoreThing(name string) {
	printHours()
	fmt.Println(fmt.Sprintf("Hello World, %s", name))
}

func printHours() {
	hours, minutes, _ := time.Now().Clock()
	fmt.Print(color.BlueString(fmt.Sprintf("%d:%02d - ", hours, minutes)))
}
