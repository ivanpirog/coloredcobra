package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "This is an example of using ColoreCobra library.",
	Long: "This is just an example of using the ColoredCobra library. \n" +
		"Project home: https://github.com/ivanpirog/coloredcobra",
	Example: "There is a simple example of the Examples section.\n" +
		"Just try commands:\n\n" +
		"example help\n" +
		"example help test",
	Aliases: []string{"alias1", "alias2", "alias3"},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No commands given. Run 'example help' for usage help.\n" +
			"Also try commands:\n\n" +
			"example help\n" +
			"example help test")
	},
}

var (
	// Used for flags.
	flag1 string
	flag2 string
	abc   []bool
)

func Execute() {

	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Aliases:       cc.Bold + cc.Italic,
		CmdShortDescr: cc.HiRed,
		Example:       cc.Italic,
		ExecName:      cc.Bold,
		Flags:         cc.Bold,
		FlagsDescr:    cc.HiRed,
		FlagsDataType: cc.Italic,
	})

	rootCmd.PersistentFlags().StringVarP(&flag1, "flag", "f", "", "some flag")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&flag2, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().BoolSliceVar(&abc, "zzz", []bool{true, false}, "usage of bools")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
