package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/mm"
	"strings"
)

func main() {
	fmt.Println("frist thing") // godoc comments

	// arrays
	names := [3]string{"a", "b", "c"}
	fmt.Println(names)

	var newvar [3]string
	newvar[0] = "a"
	newvar[1] = "b"
	newvar[2] = "c"

	fmt.Println(newvar)

	// arrays - slicing
	namesN := []string{}
	namesN = append(namesN, "a")
	namesN = append(append(namesN, "b"), "c")

	fmt.Println(names)

	// arrays - make

	certsList := make([]string, 0) // u can make an array with 0 elements.
	nn := make([]string, 4)
	nn[0] = "a"
	nn[1] = "b"
	nn[2] = "c"
	nn[3] = "d"
	nn = append(nn, "e")
	fmt.Println(nn)

	// maps
	birth := map[string]string{
		"a": "b",
		"c": "d",
	}

	fmt.Println(birth, birth["a"])

	age := map[string]int{}
	age["a"] = 11
	age["b"] = 12
	age["c"] = 13
	age["d"] = 14

	fmt.Println(age)

	// if conditional
	if age["a"] < 12 {
		fmt.Println(age["a"])
	} else if age["a"] == 12 {
		fmt.Println("no")
	} else {
		fmt.Println("no")
	}

	//switch conditional
	switch {
	case age["a"] < 12:
		fmt.Println("switch: ", age["a"])
	case age["a"] == 12:
		fmt.Println("switch no")
	case age["a"] > 12:
		fmt.Println("switch no")
	}

	// switch conditional 2
	switch age["a"] {
	case 11, 22, 12:
		fmt.Println("switch 2 : ", age["a"])
	case 223, 44:
		fmt.Println("switch 2: ", age["a"])
	}

	// for loop1
	for n, a := range age {
		fmt.Println(fmt.Sprintf("name: %s, age: %s", n, a))
	}

	// for loop2
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for loop3
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// continue and break in for loops
	aa := 0
	for aa < 10 {
		if aa%2 == 0 {
			aa++
			continue
		} else if aa == 5 {
			break
		}
		aa++
	}

	// functions
	//pp(["this is now in func"])
	fmt.Println(greeting("Chai", "Hello"))
	fmt.Println(mm.Greeting("Chai", "Hellow"))

	// read user input
	// readUserInput()

	//file operations
	fileOp()

}

func pp(ar []string) {
	for i := range ar {
		fmt.Println(i)
		fmt.Println()
	}
}

func greeting(name, greeting string) string {
	return fmt.Sprintf("%s %s", greeting, name)
}

func readUserInput() {
	//read user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("greeting: ")
	gr, _ := reader.ReadString('\n')
	gr = strings.TrimSpace(gr)
	fmt.Println("name: ")
	na, _ := reader.ReadString('\n')
	na = strings.TrimSpace(na)
	fmt.Println(mm.Greeting(na, gr))
}

func fileOp() {
	dir, _ := os.Getwd()
	fileT := dir + "\\dummy"

	f, err := os.OpenFile(fileT, os.O_WRONLY, 0644)
	if err != nil {
		f, err := os.Create(fileT)
		if err != nil {
			fmt.Println("error: unable to create and open file")
			os.Exit(2)
		} else {
			f.WriteString("write some stirng")
		}
	} else {
		f.WriteString("write some stirng")
	}

	defer f.Close()
}
