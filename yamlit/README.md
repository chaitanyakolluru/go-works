# yamlit version v0.4.0
A CLI written in GO that provides an easier to work-with command line interface to work with Terraform code. The cli can be used to prepare yaml files that define your infrastructure state and then have the cli convert yaml to Terraform code and run it to deploy resources as expected. Cli also offers ways to test existing terraform code by deploying terraform code and by using a test-conditions yaml file data, tests if the resultant objects within state file and properties are created as expected.

~~~

Yamlit is a cli that can be used to interact with AA/Terraform repo
to create, run and test terraform code.

Subcommands supported:
  - discover
  - populate
  - generate
  - plan
  - apply
  - test
  - version

Usage:
  yamlit [flags]
  yamlit [command]

Available Commands:
  apply       Subcommand that takes in variables-in-modules.yaml and runs apply to create resources.
  discover    Discover services offered within AA/Terraform.
  generate    This option take in the filled in variables-in-modules.yaml and generates a main.tf file under cwd/.temp.
  help        Help about any command
  plan        Subcommand that takes in the variables-in-modules.yaml file and returns 'terraform plan' output as stdout.
  populate    Populate yaml that generates variables-in-mouldes.yaml file.
  test        Test subcommand lets you provide a --testfile and a --testcondition file, run terraform apply and test each condition aginst the resulting terraform.tfstate file.
  version     Displays version of yamlit

Flags:
  -h, --help   help for yamlit

Use "yamlit [command] --help" for more information about a command.

~~~
