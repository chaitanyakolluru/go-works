package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	// allows u to rename the import as homedir. doing coz name has '-'
)

var name string
var greeting string
var preview bool
var prompt bool
var debug bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// all that initially went into main fucntion will go here.

		// show usage if flags are invalid
		if prompt == false && (name == "" || greeting == "") {
			cmd.Usage()
			os.Exit(1)
		}

		// optionally print flags and exit based on debug env variable
		if debug {
			fmt.Println("Name: ", name)
			fmt.Println("Greeting: ", greeting)
			fmt.Println("Prompt: ", prompt)
			fmt.Println("Preview: ", preview)
		}

		if prompt {
			name, greeting = renderPrompt()
		}

		// generate message
		m := buildMessage(name, greeting)

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

	},
}

func buildMessage(name, greeting string) string {
	return fmt.Sprintf("%s %s", greeting, name)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "name")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "greeting")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "prompt")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Preview message")

	if os.Getenv("DEBUG") != "" {
		debug = true
	}
}
