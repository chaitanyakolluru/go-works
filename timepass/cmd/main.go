package main

import (
	"fmt"
	"timepass/pkg/life"
)

func main() {
	myLife := life.CreateLife("Chaitanya Kolluru", 34, "male", "Software Engineer", 0.5, 0.4, 0.3, 0.3, 0.2, 0.7, 0.3, 0.2)
	fmt.Println(myLife)
}
