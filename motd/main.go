package main

import (
	"bufio"
	"flag"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	// define flags
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// parse flags
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "greeting phrase")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and greeting")
	flag.BoolVar(&preview, "preview", false, "use preview to op message without writing to /etc/motd")

	flag.Parse()

	// show usage if flags are invalid
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// optionally print flags and exit based on debug env variable
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name: ", name)
		fmt.Println("Greeting: ", greeting)
		fmt.Println("Prompt: ", prompt)
		fmt.Println("Preview: ", preview)
	}
	// conditionally read from stdin
	if prompt {
		name, greeting = renderPrompt()
	}

	// generate message
	m := message.Greeting3(name, greeting)

	// either preview or write it to a file
	if preview {
		fmt.Println(m)
	} else {
		//write content
		f, err := os.OpenFile("C:\\Users\\Chaitanya.Kolluru\\Desktop\\AA-new\\tuna1\\go\\src\\motd\\output", os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("error: unable to open")
			os.Exit(2)
		}

		defer f.Close() // defer lets everything in the context run and then runs what ever is in it.

		_, err = f.Write([]byte(m))
		if err != nil {
			fmt.Println("error: unable to open")
			os.Exit(3)
		}
	}

}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("your greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}
