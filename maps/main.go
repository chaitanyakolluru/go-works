package main
import "fmt"

func main() {
	birth:= map[string]string{
		"a": "b",
		"c": "d",
	}

	fmt.Println(birth,birth["a"])

	age:= map[string]int{}
	age["a"] = 11
	age["b"] = 55

	fmt.Println(age)
}