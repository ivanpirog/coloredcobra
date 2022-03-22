package coloredcobra

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "test",
		Short:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Long:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Example: "There is an example.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Test")
		},
	}

	cfg *Config
)

type colorTest struct {
	colorInt   uint8
	colorColor *color.Color
}

var colorTests = []colorTest{
	{White, color.New(color.FgWhite)},
	{Black, color.New(color.FgBlack)},
	{Red, color.New(color.FgRed)},
	{Green, color.New(color.FgGreen)},
	{Yellow, color.New(color.FgYellow)},
	{Blue, color.New(color.FgBlue)},
	{Magenta, color.New(color.FgMagenta)},
	{Cyan, color.New(color.FgCyan)},
	{HiRed, color.New(color.FgHiRed)},
	{HiGreen, color.New(color.FgHiGreen)},
	{HiYellow, color.New(color.FgHiYellow)},
	{HiBlue, color.New(color.FgHiBlue)},
	{HiMagenta, color.New(color.FgHiMagenta)},
	{HiCyan, color.New(color.FgHiCyan)},
	{HiWhite, color.New(color.FgHiWhite)},
	{White + Bold + Italic + Underline, color.New(color.FgWhite, color.Bold, color.Italic, color.Underline)},
	{Red + Bold, color.New(color.FgRed, color.Bold)},
	{Yellow + Italic, color.New(color.FgYellow, color.Italic)},
	{Blue + Underline, color.New(color.FgBlue, color.Underline)},
	{Bold, color.New(color.FgWhite, color.Bold)},
	{Italic, color.New(color.FgWhite, color.Italic)},
	{Underline, color.New(color.FgWhite, color.Underline)},
}

func TestGetColor(t *testing.T) {
	for _, test := range colorTests {
		res1 := fmt.Sprintf("%v", getColor(test.colorInt))
		res2 := fmt.Sprintf("%v", test.colorColor)
		if res1 != res2 {
			t.Errorf("got: %s, expected: %s", res1, res2)
		}
	}
}

type templateTest struct {
	in, out string
}

var templateTestHeadings = []templateTest{
	{`Usage:`, `{{HeadingStyle "Usage:"}}`},
	{`Aliases:`, `{{HeadingStyle "Aliases:"}}`},
	{`Examples:`, `{{HeadingStyle "Examples:"}}`},
	{`Available Commands:`, `{{HeadingStyle "Available Commands:"}}`},
	{`Global Flags:`, `{{HeadingStyle "Global Flags:"}}`},
	{`Additional help topics:`, `{{HeadingStyle "Additional help topics:"}}`},
	{`Flags:`, `{{HeadingStyle "Flags:"}}`},
}

var templateTests = []templateTest{
	{`{{rpad .Name .NamePadding}}`, `{{rpad (CommandStyle .Name) (sum .NamePadding 12)}}`},
	{`{{ rpad .Name .NamePadding }}`, `{{rpad (CommandStyle .Name) (sum .NamePadding 12)}}`},
	{`{{rpad .CommandPath .CommandPathPadding}}`, `{{rpad (CommandStyle .CommandPath) (sum .CommandPathPadding 12)}}`},
	{`{{   rpad .CommandPath .CommandPathPadding   }}`, `{{rpad (CommandStyle .CommandPath) (sum .CommandPathPadding 12)}}`},
	{`{{range .Commands}}{{.Short}}`, `{{range .Commands}}{{CmdShortStyle .Short}}`},
	{`{{range .Commands}}{{	.Short  }}`, `{{range .Commands}}{{CmdShortStyle .Short}}`},
	{`{{.CommandPath}}`, `{{ExecStyle .CommandPath}}`},
	{`{{ .CommandPath	}}`, `{{ExecStyle .CommandPath}}`},
	{`{{.UseLine}}`, `{{UseLineStyle .UseLine}}`},
	{`{{	.useline }}`, `{{UseLineStyle .UseLine}}`},
	{`{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}`, `{{FlagStyle .LocalFlags.FlagUsages | trimTrailingWhitespaces}}`},
	{`{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}`, `{{FlagStyle .InheritedFlags.FlagUsages | trimTrailingWhitespaces}}`},
	{`{{.NameAndAliases}}`, `{{AliasStyle .NameAndAliases}}`},
	{`{{	.NameAndAliases	}}`, `{{AliasStyle .NameAndAliases}}`},
	{`{{.Example}}`, `{{ExampleStyle .Example}}`},
	{`{{ .Example}}`, `{{ExampleStyle .Example}}`},
}

func TestTemplateReplaces(t *testing.T) {
	cfg = &Config{
		RootCmd:         rootCmd,
		Headings:        HiCyan + Bold + Underline,
		Commands:        HiYellow + Bold,
		CmdShortDescr:   HiRed,
		ExecName:        Bold,
		Flags:           Bold,
		FlagsDataType:   Italic,
		FlagsDescr:      HiRed,
		Aliases:         HiMagenta + Underline,
		Example:         Italic,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	}
	Init(cfg)
	rootCmd.PersistentFlags().StringP("flag", "f", "", "Flag description")
	rootCmd.UsageString()

	// No extra new lines
	for _, test := range templateTestHeadings {
		makeTemplateTest(test, t)
	}
	for _, test := range templateTests {
		makeTemplateTest(test, t)
	}

	// New line at the end of template
	cfg.NoBottomNewline = false
	for _, test := range templateTestHeadings {
		test.out = test.out + "\n"
		makeTemplateTest(test, t)
	}
	for _, test := range templateTests {
		test.out = test.out + "\n"
		makeTemplateTest(test, t)
	}

	// Extra new lines before and after headings
	cfg.NoBottomNewline = true
	cfg.NoExtraNewlines = false
	for _, test := range templateTestHeadings {
		test.out = "\n" + test.out + "\n"
		makeTemplateTest(test, t)
	}

}

func makeTemplateTest(test templateTest, t *testing.T) {
	rootCmd.SetUsageTemplate(test.in)
	Init(cfg)
	res := rootCmd.UsageTemplate()
	if res != test.out {
		t.Errorf("got: %v, expected: %v", res, test.out)
	}
}
