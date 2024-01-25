package main
import "fmt"

func main() {
	age:= map[string]int{
		"a": 55,
		"c": 67,
		"b": 12,
	}

	fmt.Println("first for")
	for name,age := range age { // range only works as an expression in the context of a for loop
		fmt.Println("name: ",name,"age: ",age)
		fmt.Println(fmt.Sprintf("name: %s age: %s", name,age))
	}
	fmt.Println("second for")
	for i:=1; i<=10; i++ {
		fmt.Println(i)
	}

	fmt.Println("third for")
	a:=0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("fourth for")
	aa:=0
	for aa < 10 {
		if aa%2 == 0 {
			aa++
			continue
		} else if aa == 5 {
			break
		}

		fmt.Println(aa)
		aa++
	}
}