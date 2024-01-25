package main

import (
	"bufio"
	"flag"
	"fmt"
	"motd-practice/message"
	"os"
	"strings"
)

func main() {
	var name string
	var greeting string
	var prompt bool
	var preview bool

	flag.StringVar(&name, "name", "", "name to use")
	flag.StringVar(&greeting, "greeting", "", "greeting phrase")
	flag.BoolVar(&prompt, "prompt", false, "use prompt or no")
	flag.BoolVar(&preview, "preview", false, "use preview or no")

	flag.Parse()

	if !prompt && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	_, debug_present := os.LookupEnv(("DEBUG"))
	if debug_present {
		fmt.Println("Name: ", name)
		fmt.Println("Greeting: ", greeting)
		fmt.Println("Prompt: ", prompt)
		fmt.Println("Preview: ", preview)
	}

	if prompt {
		name, greeting = renderPrompt()
	}

	mess := message.Greeting(name, greeting)

	if preview {
		fmt.Println(mess)
	} else {
		f, err := os.OpenFile("./file", os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("unable to open file")
			os.Exit(2)
		}
		defer f.Close()

		_, err = f.Write([]byte(mess))
		if err != nil {
			fmt.Println("unable to write")
			os.Exit(3)
		}
	}
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("your greeting...")
	greeting, _ = reader.ReadString(('\n'))
	greeting = strings.TrimSpace(greeting)

	fmt.Println("your name...")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace((name))
	return
}

// motd-practice ➤ go run main.go -name chai -greeting howdy!
// motd-practice ➤ go run main.go -name chai -greeting howdy!
// motd-practice ➤ go run main.go -name chai -greeting howdy! -preview
// howdy! chai
// motd-practice ➤ go run main.go -name chai -greeting howdy! -preview
// motd-practice ➤ go run main.go -prompt
// your greeting...
// howdy!
// your name...
// chai
// motd-practice ➤
