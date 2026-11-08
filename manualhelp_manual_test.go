//go:build tinygo || js

package coloredcobra

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestManualHelpSections(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	root := &cobra.Command{Use: "app", Run: func(*cobra.Command, []string) {}}
	sub := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Example: "  app serve -a :8080",
		Short:   "serve it",
		Run:     func(*cobra.Command, []string) {},
	}
	root.AddCommand(sub)
	Init(&Config{RootCmd: root, Headings: HiBlue})

	var b bytes.Buffer
	sub.SetOut(&b)
	sub.SetErr(&b)
	if err := sub.Usage(); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Usage:", "Aliases:\n  serve, s", "Examples:\n  app serve -a :8080"} {
		if !strings.Contains(b.String(), want) {
			t.Errorf("usage lacks %q:\n%s", want, b.String())
		}
	}
}

func TestManualHelpLayout(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	root := &cobra.Command{Use: "app", Long: "app does things.", Run: func(*cobra.Command, []string) {}}
	Init(&Config{RootCmd: root, Headings: HiBlue})

	var out, errOut bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&errOut)
	root.SetArgs([]string{"--help"})
	if err := root.Execute(); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out.String(), "app does things.\n\nUsage:\n") {
		t.Errorf("help on stdout = %q, want the description, a blank line, then Usage", out.String())
	}
	if errOut.Len() != 0 {
		t.Errorf("help wrote to stderr: %q", errOut.String())
	}
}
