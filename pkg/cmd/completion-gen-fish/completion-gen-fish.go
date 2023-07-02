package completiongenfish

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "completion-gen-fish",
		Short: "Generate a fish-shell completion file from any CLI tool help or manual",
	}
	target, reference string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&target, "target-command", "t", "")
	rootCmd.PersistentFlags().StringVar(&reference, "reference", "r", "")
}
