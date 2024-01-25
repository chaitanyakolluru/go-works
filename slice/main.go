package main
import "fmt"

func main () {
	names:= []string{}
	names= append(names,"a")
	names= append(names,"b", "c", "d")

	fmt.Println(names)
}