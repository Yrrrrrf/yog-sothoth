package hack

import (
	"fmt"
	"src/hack_go"
)

// RunTool acts as a thin wrapper sending user commands to the educational library.
func RunTool(toolName string) (string, error) {
	if toolName == "" {
		return "", fmt.Errorf("tool name is required")
	}

	// For now, regardless of the tool, we will run the underlying DoHackGoStuff wrapper
	// In the future this can be expanded to dispatch based on toolName.
	message := hack_go.DoHackGoStuff()
	return fmt.Sprintf("Hack API response for '%s':\n%s", toolName, message), nil
}
