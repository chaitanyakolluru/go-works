package main

import "fmt"

func main() {
	i := make(map[interface{}]interface{})
	t := make(map[interface{}]interface{})
	t[2] = 3
	i[1] = t

	for k, v := range i {
		fmt.Println(k)

		for ki, vi := range v.(map[interface{}]interface{}) {

			fmt.Println(ki, vi)
			fmt.Printf("%v %T", vi, vi)
		}
	}

}
