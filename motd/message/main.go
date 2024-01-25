package message
import "fmt"

func Greeting3(name string,message string) (salutation string) { // salutation automatically gets initialized
	salutation = fmt.Sprintf("%s %s",message,name)
	return
}