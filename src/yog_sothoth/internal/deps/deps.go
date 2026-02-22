package deps

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"src/yog_sothoth/pkg/ui"
	"src/yog_sothoth/scripts"
)

func detectRuntime() string {
	if _, err := os.Stat("deno.json"); err == nil {
		return "deno"
	}
	if _, err := os.Stat("deno.jsonc"); err == nil {
		return "deno"
	}
	if _, err := os.Stat("bunfig.toml"); err == nil {
		return "bun"
	}
	if _, err := os.Stat("bun.lockb"); err == nil {
		return "bun"
	}
	
	// Fallback detection
	if _, err := os.Stat("package.json"); err == nil {
		fmt.Println(ui.RenderWarn("Node.js is never supported. Defaulting to Bun as runtime proxy for package.json..."))
		return "bun"
	}

	return ""
}

// Reborn runs the clean up and reinstall
func Reborn(runtime string, deep, dryRun, noInstall bool) error {
	if runtime == "" {
		runtime = detectRuntime()
		if runtime == "" {
			return fmt.Errorf("could not auto-detect runtime (deno or bun). Please specify it")
		}
		fmt.Println(ui.RenderInfo(fmt.Sprintf("Auto-detected runtime: %s", runtime)))
	}

	if runtime != "deno" && runtime != "bun" {
		return fmt.Errorf("unsupported runtime: %s (Node.js is never supported)", runtime)
	}

	scriptName := fmt.Sprintf("reborn_%s.sh", runtime)
	scriptContent, err := scripts.Files.ReadFile(scriptName)
	if err != nil {
		return fmt.Errorf("internal error: could not load script %s: %w", scriptName, err)
	}

	// Write script to a temporary file to execute
	tmpDir := os.TempDir()
	tmpScriptPath := filepath.Join(tmpDir, scriptName)
	if err := os.WriteFile(tmpScriptPath, scriptContent, 0700); err != nil {
		return fmt.Errorf("failed to write tmp script: %w", err)
	}
	defer os.Remove(tmpScriptPath)

	deepStr := "false"
	if deep {
		deepStr = "true"
	}
	dryRunStr := "false"
	if dryRun {
		dryRunStr = "true"
	}
	noInstallStr := "false"
	if noInstall {
		noInstallStr = "true"
	}

	cmd := exec.Command("bash", tmpScriptPath, deepStr, dryRunStr, noInstallStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("reborn script failed: %w", err)
	}

	return nil
}
