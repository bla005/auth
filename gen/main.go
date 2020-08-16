package main

import (
	"os"

	"github.com/spf13/cobra"
)

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "update",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := os.Create("jet.go")
			if err != nil {
				return err
			}
			defer f.Close()
			f.Write([]byte(`hi nigga`))
			return nil
		},
	}
}

func main() {
	cmd := &cobra.Command{
		Use:          "wtf",
		Short:        "hello",
		SilenceUsage: true,
	}

	cmd.AddCommand(printTimeCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
