package mm

import "fmt"

func Greeting(name, greeting string) string {
	return fmt.Sprintf("%s %s", greeting, name)
}
