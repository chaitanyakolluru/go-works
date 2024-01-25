package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputyaml, environment, location, username, token string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yamlit",
	Short: "Tool used to generate and run terraform code by interfacing with modules within AA/Terraform",
	Long: `Yamlit is a cli that can be used to interact with AA/Terraform repo
to create, run and test terraform code. 
	
Subcommands supported:
- discover
- populate
- generate
- plan
- apply
- test 
- version
`,

	Run: func(cmd *cobra.Command, args []string) {

		cmd.Usage()
		os.Exit(1)

		// azure sdk for authentication and terraform command execution

	},
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
	rootCmd.PersistentFlags().StringVarP(&username, "username", "", "", `Optional flag that lets you provide username to authenticate to ghe.aa.com.
Alternatively, you can set environmental variable GHE_USERNAME`)
	rootCmd.PersistentFlags().StringVarP(&token, "token", "", "", `Optional flag that lets you provide token to authenticate to ghe.aa.com.
Alternatively, you can set environmental variable GHE_TOKEN`)
}
