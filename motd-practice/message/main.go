package message

import "fmt"

func Greeting(name string, message string) (salutation string) {
	salutation = fmt.Sprintf("%s %s", message, name)
	return
}
