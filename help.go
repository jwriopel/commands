package commands

import (
	"fmt"
	"os"
)

// HelpCommand is an example command that will print the Short description
// for all registered commands.
var HelpCommand = &Command{
	Name:  "help",
	Usage: "help",
	Short: "List available commands.",
}

func runHelp(c *Command, args []string) {
	for _, cmd := range commands {
		fmt.Fprintf(os.Stderr, "\t%s - %s\n", cmd.Name, cmd.Short)
	}
	fmt.Fprintln(os.Stderr, "\nUse <command -help> for detailed help.")
}

func init() {
	HelpCommand.Run = runHelp
}
