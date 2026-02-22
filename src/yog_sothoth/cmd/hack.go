package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"src/yog_sothoth/internal/hack"
	"src/yog_sothoth/pkg/ui"
)

var hackCmd = &cobra.Command{
	Use:   "hack [tool]",
	Short: "Cybersecurity Learning Tools",
	Long:  "Exposes the hack_go library commands for security research, network analysis, and educational purposes. All tools are strictly for authorized use only.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		
		result, err := hack.RunTool(toolName)
		if err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			os.Exit(1)
		}

		fmt.Println(ui.RenderTitle("Yog-Sothoth Security Utils"))
		fmt.Println(ui.RenderBox(result))
	},
}

func init() {
	rootCmd.AddCommand(hackCmd)
}
