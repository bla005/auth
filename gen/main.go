package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "curtime",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			cmd.Println("whatever", now)
			return nil
		},
	}
}

func main() {
	cmd := &cobra.Command{
		Use:          "gifm",
		Short:        "hello",
		SilenceUsage: true,
	}

	cmd.AddCommand(printTimeCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
