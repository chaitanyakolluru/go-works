package main
import "fmt"

func main() {
	names:= [3]string{"a","b","c"}
	fmt.Println(names)

	var newvar [3]string
	newvar[0] = "a"
	newvar[1] = "b"
	newvar[2] = "c"

	fmt.Println(newvar)
}