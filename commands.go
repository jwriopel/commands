package commands

// The commands package provides a simple framework for building commands
// and sub-commands. Heavily inspired by https://golang.org/src/cmd/go/internal/base/base.go

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Command types are used for building subcommands.
type Command struct {
	Name  string
	Usage string
	Short string
	Flags flag.FlagSet
	Run   func(*Command, []string)
}

// Contains all available commands.
var commands map[string]*Command

func init() {
	commands = make(map[string]*Command)
	Add(HelpCommand)
}

// Add a Command.
func Add(c *Command) {
	if _, ok := commands[c.Name]; ok {
		fmt.Fprintf(os.Stderr, "over writing existing command: %s\n", c.Name)
	}
	commands[c.Name] = c
}

// Get looks up a command by name.
func Get(commandName string) *Command {
	command, ok := commands[commandName]
	if !ok {
		return nil
	}
	return command
}

// Run will run execute the command contained in cmdLine.
func Run(cmdLine string) error {
	cmdLine = strings.Trim(cmdLine, " \n")
	cmdParts := strings.Split(cmdLine, " ")
	cmdName := cmdParts[0]

	cmd := Get(cmdName)
	if cmd == nil {
		return fmt.Errorf("%s - command not found", cmdName)
	}

	doRun := true

	cmd.Flags.Usage = func() {
		fmt.Fprintf(os.Stderr, cmd.Usage)
		cmd.Flags.PrintDefaults()
		fmt.Print("\n")
		doRun = false
	}
	cmd.Flags.Parse(cmdParts[1:])
	if doRun {
		cmd.Run(cmd, cmd.Flags.Args())
	}
	return nil
}
