package hello

import (
	"fmt"
	"time"
)

// SayHelloAndOneMoreThing function: Says hello
func SayHelloAndOneMoreThing(name string) {
	printHours()
	fmt.Println(fmt.Sprintf("Hello World, %s", name))
}

func printHours() {
	hours, minutes, _ := time.Now().Clock()
	fmt.Println(fmt.Sprintf("%d:%02d ", hours, minutes))
}
