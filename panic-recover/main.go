package main

import "fmt"

func main() {
	var action int
	fmt.Println("input 1 or 2")
	fmt.Scanln(&action)
	switch action {
	case 1:
		fmt.Println("student")
	case 2:
		fmt.Println("professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))

	}

	defer func() {
		action := recover()
		fmt.Println(action)
	}()
}
