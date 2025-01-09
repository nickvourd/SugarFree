package Arguments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// infoArgument represents the 'info' command in the CLI.
var infoArgument = &cobra.Command{
	// Use defines how the command should be called.
	Use:          "info",
	Short:        "info command",
	SilenceUsage: true,
	Aliases:      []string{"INFO", "Info"},

	// RunE defines the function to run when the command is executed.
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.New(os.Stderr, "[!] ", 0)

		// Call function named ShowAscii
		ShowAscii()

		// Check if additional arguments were provided.
		if len(os.Args) <= 2 {
			// Show help message.
			err := cmd.Help()
			if err != nil {
				logger.Fatal("Error ", err)
				return err
			}

			// Exit the program.
			os.Exit(0)
		}

		// Get variables from the command line.
		file, _ := cmd.Flags().GetString("file")
		output, _ := cmd.Flags().GetString("output")

		// Check if the file flag is empty.
		if file == "" {
			logger.Fatal("Error: Input file is missing. Please provide it to continue...\n")
		}

		fmt.Println("File: ", file)
		fmt.Println("Output: ", output)

		return nil
	},
}
