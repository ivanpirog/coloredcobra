package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Host string
	Port int
)

var helloCmd = &cobra.Command{
	Use:   "test",
	Short: "Test command for clarity",
	Long:  `This command prints a simple message "Hello world!"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func init() {
	helloCmd.Flags().StringVarP(&Host, "host", "s", "localhost", "A host for the Hello World server")
	helloCmd.Flags().IntVarP(&Port, "port", "p", 8080, "A port for the Hello World server")
	rootCmd.AddCommand(helloCmd)
}
