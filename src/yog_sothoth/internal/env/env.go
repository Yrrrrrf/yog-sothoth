package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	"src/yog_sothoth/pkg/ui"
)

// Load traverses up the tree to find and load a .env file
func Load(showValues bool) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			fmt.Println(ui.RenderInfo(fmt.Sprintf("Loading .env from %s", envPath)))
			
			envMap, err := godotenv.Read(envPath)
			if err != nil {
				return err
			}
			
			err = godotenv.Load(envPath)
			if err != nil {
				return err
			}
			
			count := 0
			for k, v := range envMap {
				displayStr := k
				if showValues {
					displayStr = fmt.Sprintf("%s=%s", k, v)
				}
				fmt.Println(ui.RenderSuccess(fmt.Sprintf("Loaded: %s", displayStr)))
				count++
			}
			
			fmt.Println(ui.RenderSuccess(fmt.Sprintf("Finished: Loaded %d variables from %s.", count, envPath)))
			return nil
		}

		parent := filepath.Dir(dir)
		if parent == dir { // Reached root
			break
		}
		dir = parent
	}

	return fmt.Errorf("no .env file found in directory tree")
}

// Check validates that all variables in .env.example are present in .env
func Check() error {
	envVars, err := godotenv.Read(".env")
	if err != nil {
		return fmt.Errorf("failed to read .env: %w", err)
	}

	exampleVars, err := godotenv.Read(".env.example")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(ui.RenderInfo("No .env.example found. Nothing to check."))
			return nil
		}
		return fmt.Errorf("failed to read .env.example: %w", err)
	}

	missing := []string{}
	for key := range exampleVars {
		if _, ok := envVars[key]; !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		errMsg := "Missing variables in .env:\n"
		for _, m := range missing {
			errMsg += ui.RenderItem(m) + "\n"
		}
		return fmt.Errorf(errMsg)
	}

	fmt.Println(ui.RenderSuccess("All variables from .env.example are present in .env!"))
	return nil
}

// Diff shows variables in .env that aren't in .env.example and vice-versa
func Diff() error {
	envVars, err := godotenv.Read(".env")
	if err != nil {
		return fmt.Errorf("failed to read .env: %w", err)
	}

	exampleVars, err := godotenv.Read(".env.example")
	if err != nil {
		return fmt.Errorf("failed to read .env.example: %w", err)
	}

	extra := []string{}
	for key := range envVars {
		if _, ok := exampleVars[key]; !ok {
			extra = append(extra, key)
		}
	}

	missing := []string{}
	for key := range exampleVars {
		if _, ok := envVars[key]; !ok {
			missing = append(missing, key)
		}
	}

	hasDrift := len(extra) > 0
	hasMissing := len(missing) > 0

	if hasDrift {
		fmt.Println(ui.RenderWarn("Variables in .env but missing from .env.example (potential drift):"))
		for _, e := range extra {
			fmt.Println(ui.RenderItem(e))
		}
	}

	if hasMissing {
		fmt.Println()
		fmt.Println(ui.RenderError("Variables in .env.example but missing from .env:"))
		for _, m := range missing {
			fmt.Println(ui.RenderItem(m))
		}
	}

	if !hasDrift && !hasMissing {
		fmt.Println(ui.RenderSuccess("No drift detected! .env and .env.example are perfectly synced."))
	}

	return nil
}

// Sync prompts interactively for any missing values from .env.example and appends them to .env
func Sync() error {
	envVars, err := godotenv.Read(".env")
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read .env: %w", err)
	}
	if envVars == nil {
		envVars = make(map[string]string)
	}

	exampleVars, err := godotenv.Read(".env.example")
	if err != nil {
		return fmt.Errorf("failed to read .env.example: %w", err)
	}

	missing := []string{}
	for key := range exampleVars {
		if _, ok := envVars[key]; !ok {
			missing = append(missing, string(key))
		}
	}

	if len(missing) == 0 {
		fmt.Println(ui.RenderSuccess(".env is fully synced with .env.example."))
		return nil
	}

	fmt.Println(ui.RenderInfo("Syncing missing variables..."))
	fmt.Println(ui.SuccessStyle.Render("Enter values:"))
	
	reader := bufio.NewReader(os.Stdin)

	f, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("could not open .env for writing: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString("\n# Added by yog env sync\n"); err != nil {
		return err
	}

	// Inline style for the prompt key without MarginBottom from ui.TitleStyle
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(ui.PrimaryColor).PaddingLeft(1)

	for _, key := range missing {
		exampleVal := exampleVars[key]
		
		// The prompt prints the Key, then the example dimmed.
		// We do NOT print a newline so the user types directly on this line.
		fmt.Printf("%s %s", 
			keyStyle.Render(key+":"),
			lipgloss.NewStyle().Foreground(ui.MutedColor).Render(exampleVal),
		)
		
		// Move the cursor back by the length of the example text 
		// so the user types "over" the example text (creating the ghost effect).
		if len(exampleVal) > 0 {
			offset := lipgloss.Width(keyStyle.Render(key+":")) + 1 // +1 for the space after the colon
			fmt.Printf("\r\033[%dC", offset)
		}

		val, _ := reader.ReadString('\n')
		val = strings.TrimSpace(val)

		// If user just hit enter, fall back to the example
		if val == "" {
			val = exampleVal
		}

		line := fmt.Sprintf("%s=\"%s\"\n", key, val)
		if _, err := f.WriteString(line); err != nil {
			return fmt.Errorf("failed to write to .env: %w", err)
		}
	}

	fmt.Println(ui.RenderSuccess("Sync complete! Added missing variables to .env"))
	return nil
}
