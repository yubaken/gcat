package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Globals
var (
	version bool
)

var Root = &cobra.Command{
	Use:   "gcat",
	Short: "short description",
	Long:  `long description`,
}

func init() {
	Root.Run = Run
	Root.Flags().BoolVarP(&version, "version", "V", false, "Print the version number")
	cobra.OnInitialize()
}

// CheckArgs checks there are enough arguments and prints a message if not
func CheckArgs(MinArgs, MaxArgs int, cmd *cobra.Command, args []string) {
	if -1 != MinArgs && len(args) < MinArgs {
		_ = cmd.Usage()
		fmt.Fprintf(os.Stderr, "Command %s needs %d arguments mininum\n", cmd.Name(), MinArgs)
		os.Exit(1)
	} else if -1 != MaxArgs && len(args) > MaxArgs {
		_ = cmd.Usage()
		fmt.Fprintf(os.Stderr, "Command %s needs %d arguments maximum\n", cmd.Name(), MaxArgs)
		os.Exit(1)
	}
}

func Run(cmd *cobra.Command, args []string) {
	if version {
		// Version()
		os.Exit(0)
	}

	CheckArgs(1, -1, cmd, args)

	var file *os.File
	var err error

	var size = len(args)
	for i := 0; i < size; i++ {
		var name string = args[i]

		file, err = os.Open(name)
		if err != nil {
			fmt.Println("gcat: file not found : " + name)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	os.Exit(0)
}
