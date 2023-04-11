package main

import (
	"fmt"
	"os"

	cmd "github.com/SuperJourney/gopen/cmd/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: ""}
	rootCmd.AddCommand(cmd.MigrationCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
