package main
import "fmt"

func main() {
	fmt.Println(greeting("chai","hello"))	
	fmt.Println(greeting2("chai","hello"))	
}

func greeting(name string, message string) string {
	return fmt.Sprintf("%s %s",message,name)
}
/*
func greeting(name,message string) string {
	return fmt.Sprintf("%s %s",message,name)
}
*/

func greeting2(name string,message string) (salutation string) { // salutation automatically gets initialized
	salutation = fmt.Sprintf("%s %s",message,name)
	return
}

