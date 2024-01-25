package main
import "fmt"

func main() {
	age:= map[string]int{
		"a": 55,
		"c": 67,
	}

	if age["a"] < 65 {
		fmt.Println("not retiring yet")
	} else if age["a"] == 65 {
		fmt.Println("lucky 65")
	} else {
		fmt.Println("retiring now")
	}


	switch {
	case age["b"] < 65:
		fmt.Println("switch: not retiring yet")
	case age["b"] == 65:
		fmt.Println("switch: lucky 65")
	case age["b"] > 65:
		fmt.Println("switch: retiring now")
	}

	switch age["a"] {
	case 55,67,65:
		fmt.Println("number matched")
	default:
		fmt.Println("not matched")
	}
}