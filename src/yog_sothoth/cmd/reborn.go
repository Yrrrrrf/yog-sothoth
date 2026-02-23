package cmd

import (
	"fmt"
	"os"

	"src/yog_sothoth/internal/deps"
	"src/yog_sothoth/pkg/ui"

	"github.com/spf13/cobra"
)

var (
	rebornDeep      bool
	rebornDryRun    bool
	rebornNoInstall bool
	rebornFull      bool
)

var rebornCmd = &cobra.Command{
	Use:   "reborn [runtime]",
	Short: "Deep-clean build artifacts and reinstall dependencies",
	Long:  `Performs a deep clean of build artifacts (.svelte-kit, node_modules, etc) and reinstalls dependencies from scratch. Auto-detects runtime (Deno or Bun) if not specified. Note: Node.js is never supported.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runtime := ""
		if len(args) > 0 {
			runtime = args[0]
		}

		if err := deps.Reborn(runtime, rebornDeep, rebornDryRun, rebornNoInstall, rebornFull); err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	rebornCmd.Flags().BoolVar(&rebornDeep, "deep", false, "also removes lock files")
	rebornCmd.Flags().BoolVar(&rebornDryRun, "dry-run", false, "shows what would be deleted without doing it")
	rebornCmd.Flags().BoolVar(&rebornNoInstall, "no-install", false, "just clean, don't reinstall")
	rebornCmd.Flags().BoolVar(&rebornFull, "full", false, "recursively clean artifacts in all subdirectories")

	rootCmd.AddCommand(rebornCmd)
}
