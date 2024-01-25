package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"practice/mm"
	"strings"
)

func main() {
	// define flags
	var name, greeting string
	var prompt, preview bool

	// parse flags
	flag.StringVar(&name, "name", "", "Name to use for the message")
	flag.StringVar(&greeting, "greeting", "", "greeting phrase")
	flag.BoolVar(&prompt, "prompt", true, "use prompt to input name and greeting")
	flag.BoolVar(&preview, "preview", true, "Preview message without writing")

	flag.Parse()
	// show usage if flags are invalid
	if name == "" || greeting == "" {
		flag.Usage()
		os.Exit(1)
	}

	// optionally print flags and exit based on debug env variable
	if os.Getenv("DEBUG") != "" {
		fmt.Println(name, greeting)
	}

	// conditionally read from stdin
	if name == "" || greeting == "" {
		name, greeting = renderPrompt()
	}

	// generate message
	mess := mm.Greeting(name, greeting)

	// op section
	fmt.Println(mess)
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)
	return
}
