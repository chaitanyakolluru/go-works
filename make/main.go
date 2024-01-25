package main
import "fmt"

func main() {
	newvar:= make([]string,4)
	newvar[0] = "a"
	newvar[1] = "b"
	newvar[2] = "b"
	newvar[3] = "c"
	newvar = append(newvar,"vv")
	fmt.Println(newvar)
}