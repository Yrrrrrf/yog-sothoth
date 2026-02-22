package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"src/yog_sothoth/internal/env"
	"src/yog_sothoth/pkg/ui"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Environment variable management",
	Long:  "Manages .env files with validation, diffing, and interactive syncing.",
}

var showValues bool

var envLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Finds and loads a .env file by walking up the directory tree",
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Load(showValues); err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}
	},
}

var envCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Validates that all variables in .env.example are present in .env",
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Check(); err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}
	},
}

var envDiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Shows variables in .env that aren't in .env.example",
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Diff(); err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}
	},
}

var envSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Interactive mode: prompts for missing values to sync .env with .env.example",
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Sync(); err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	envLoadCmd.Flags().BoolVarP(&showValues, "show-values", "s", false, "Display the loaded values alongside their keys")
	envCmd.AddCommand(envLoadCmd)
	envCmd.AddCommand(envCheckCmd)
	envCmd.AddCommand(envDiffCmd)
	envCmd.AddCommand(envSyncCmd)
	rootCmd.AddCommand(envCmd)
}
