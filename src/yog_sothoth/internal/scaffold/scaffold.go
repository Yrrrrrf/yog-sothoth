package scaffold

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"src/yog_sothoth/internal/deps"
	"src/yog_sothoth/pkg/ui"
)

func InitProject(templateName string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	templateDir := filepath.Join(home, ".config", "yog", "templates", templateName)

	if stat, err := os.Stat(templateDir); os.IsNotExist(err) || !stat.IsDir() {
		return fmt.Errorf("template '%s' not found at %s", templateName, templateDir)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	fmt.Println(ui.RenderInfo(fmt.Sprintf("Scaffolding new project using template '%s'...", templateName)))

	if err := copyDir(templateDir, cwd); err != nil {
		return fmt.Errorf("failed to copy template files: %w", err)
	}

	fmt.Println(ui.RenderSuccess("Project scaffolded successfully!"))

	// Trigger reborn based on template name (assuming template name corresponds to runtime for now)
	// We'll pass false for deep, dryRun, and noInstall
	if err := deps.Reborn(templateName, false, false, false); err != nil {
		return fmt.Errorf("post-init reborn hook failed: %w", err)
	}

	return nil
}

// copyDir recursively copies a directory tree, replacing placeholder strings later if needed.
func copyDir(src string, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return err
			}
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Ensure the destination directory exists (though copyDir handles this mostly)
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync()
}
