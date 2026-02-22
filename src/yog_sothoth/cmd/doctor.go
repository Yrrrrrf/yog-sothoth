package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"src/yog_sothoth/internal/health"
	"src/yog_sothoth/pkg/ui"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run a comprehensive health check on the project environment",
	Long:  `Runs a comprehensive health check verifying runtimes, tools, .env completeness, and basic repository files, producing a clean pass/warn/fail report.`,
	Run: func(cmd *cobra.Command, args []string) {
		results := health.RunAllChecks()

		fmt.Println(ui.RenderTitle("Yog-Sothoth Diagnostics Report"))
		fmt.Println()

		hasFailures := false

		for _, res := range results {
			var icon, text string
			
			switch res.Status {
			case "pass":
				icon = ui.SuccessStyle.Render("✓")
				text = fmt.Sprintf("%-20s : %s", res.Name, res.Message)
			case "warn":
				icon = ui.WarnStyle.Render("!")
				text = ui.WarnStyle.Render(fmt.Sprintf("%-20s : %s", res.Name, res.Message))
			case "fail":
				hasFailures = true
				icon = ui.ErrorStyle.Render("✗")
				text = ui.ErrorStyle.Render(fmt.Sprintf("%-20s : %s", res.Name, res.Message))
			}

			fmt.Printf(" %s  %s\n", icon, text)
		}

		fmt.Println()
		if hasFailures {
			fmt.Println(ui.RenderError("Environment check failed. Please fix the items above."))
			os.Exit(1)
		} else {
			fmt.Println(ui.RenderSuccess("Environment is healthy!"))
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
