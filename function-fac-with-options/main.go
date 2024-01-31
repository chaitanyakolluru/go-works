package main

import "fmt"

type item struct {
	itemNo int
}

type typeFunc func(i *item)

func printitem() typeFunc {
	return func(i *item) {
		fmt.Println("this one", i.itemNo)
	}
}

func removeitem() typeFunc {
	return func(i *item) {
		i.itemNo = 0
	}
}

func additem(data int) typeFunc {
	return func(i *item) {
		i.itemNo = data
	}
}

func main() {
	fmt.Println("Hello, World!")
	i := &item{}

	for _, f := range []typeFunc{
		additem(10),
		printitem(),
		removeitem(),
		printitem(),
	} {
		f(i)
	}
}
