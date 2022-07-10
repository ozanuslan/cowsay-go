package main

import (
	"cowsay/cow"
	"fmt"
	"os"
)

func main() {
	say := "Hello, World!"
	if len(os.Args) > 1 {say = os.Args[1]}
	fmt.Println(cow.Say(say))
}
