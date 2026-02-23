package cmd

import (
	"fmt"
	"os"

	"src/yog_sothoth/pkg/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "yog",
	Version: "0.0.1",
	Short:   "Yog-Sothoth: The Key and the Gate",
	Long:    `Yog-Sothoth is the infrastructure layer. It prepares the universe for your projects and ensures everything has a solid foundation.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}
