package cmd

import (
	"bufio"
	"fmt"
	"os"
	"practice/mm"
	"strings"

	"github.com/spf13/cobra"
)

//define flags
var name, greeting string
var prompt, preview bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-cli-practice",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		// show usage if flags are invalid
		if name == "" || greeting == "" {
			cmd.Usage()
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
	},
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
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", true, "prompt")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", true, "Preview message")

}
