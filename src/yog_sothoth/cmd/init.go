package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"src/yog_sothoth/internal/scaffold"
	"src/yog_sothoth/pkg/ui"
)

var initCmd = &cobra.Command{
	Use:   "init [template]",
	Short: "Scaffold a new project from your personal templates",
	Long: `Scaffolds a new project from templates in ~/.config/yog/templates/.
Supports options like deno, bun, svelte, go, rust, python, fullstack.
It copies the files to the current directory and runs the appropriate reborn variant.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]
		
		if err := scaffold.InitProject(templateName); err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
