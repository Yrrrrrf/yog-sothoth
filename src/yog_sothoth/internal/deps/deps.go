package deps

import (
	"fmt"
	"os"
	"os/exec"
	"src/yog_sothoth/pkg/ui"
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

	targets := []string{"node_modules", ".svelte-kit", ".vite", "dist"}

	if deep {
		if runtime == "deno" {
			targets = append(targets, "deno.lock")
		} else if runtime == "bun" {
			targets = append(targets, "bun.lockb", "bun.lock")
		}
	}

	fmt.Println(ui.RenderInfo(fmt.Sprintf("Purging %s build artifacts...", runtime)))

	for _, target := range targets {
		if _, err := os.Stat(target); err == nil {
			if dryRun {
				fmt.Println(ui.RenderWarn(fmt.Sprintf("[DRY-RUN] Would remove: %s", target)))
			} else {
				fmt.Printf("Removing: %s\n", target)
				if err := os.RemoveAll(target); err != nil {
					return fmt.Errorf("failed to remove %s: %w", target, err)
				}
			}
		}
	}

	if noInstall {
		fmt.Println(ui.RenderInfo("Skipping installation as requested."))
		return nil
	}

	if dryRun {
		fmt.Println(ui.RenderWarn(fmt.Sprintf("[DRY-RUN] Would run: %s install", runtime)))
	} else {
		fmt.Println(ui.RenderInfo("Reinstalling dependencies..."))
		cmd := exec.Command(runtime, "install")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("reborn installation failed: %w", err)
		}
	}

	fmt.Println(ui.RenderSuccess(fmt.Sprintf("%s reborn complete.", runtime)))
	return nil
}
