# Yog-Sothoth Architecture & Testing Guide

This guide provides an overview of the **Yog-Sothoth** implementation and instructions on how to manually test its components.

## Architecture Overview

Yog-Sothoth is built in Go as an infrastructure CLI tool. It strictly separates concerns into specific layers:

- **CLI Layer (`cmd/`)**: Built using [Cobra](https://github.com/spf13/cobra). Maps user inputs, flags, and arguments. It delegates all actual work to the `internal/` packages.
- **Logic Layer (`internal/`)**: Contains the business logic for each command (`deps`, `env`, `hack`, `health`, `scaffold`). This code is completely decoupled from Cobra and CLI concerns.
- **Shared Layer (`pkg/`)**: Contains configuration parsing via [Viper](https://github.com/spf13/viper) (`pkg/config`) and common UI styling using [Lip Gloss](https://github.com/charmbracelet/lipgloss) (`pkg/ui`).
- **Scripts (`scripts/`)**: Standalone bash scripts for cleaning dependencies and interacting natively with the filesystem (e.g., `reborn_deno.sh`). These are embedded directly into the Go binary using `//go:embed`.

---

## How to Test the CLI

You can test Yog-Sothoth by building the binary and invoking the various commands.

### 1. Build and Install

To test the tool across your entire system, compile it and place the binary in your Go `bin` path:

```bash
go build -o ~/go/bin/yog ./src/yog_sothoth
```

Ensure `~/go/bin` is added to your shell's `$PATH`.

### 2. Testing `yog doctor`

This command runs system health checks.

```bash
yog doctor
```
**Expected Behavior**:
- Outputs a pass/warn/fail report formatting cleanly with Lip Gloss.
- Identifies `Deno` and `Bun` runtimes natively.
- Warns if `.env` or repository files (`README.md`, `LICENSE`, `.git/`) are missing.

### 3. Testing `yog env`

This command manages `.env` variables safely and ensures no drift happens.

**Setup**:
1. Create a mock `.env.example` in an empty folder:
   ```env
   DATABASE_URL=
   API_SECRET=
   ```
2. Create a mock `.env`:
   ```env
   DATABASE_URL=postgres://localhost
   ```

**Commands to test**:

```bash
# Checks if variables in .example are missing from .env
yog env check

# Shows variables present in .env but NOT in .example (drift)
yog env diff

# Prompts you interactively to fill missing variables from .example 
yog env sync
```

### 4. Testing `yog reborn`

This command performs a deep clean of build artifacts and reinstalls dependencies. 

**Setup**:
1. Create a dummy `deno.json` or `bunfig.toml` inside a directory.
2. Create dummy cache folders: `mkdir .vite dist node_modules`

**Commands to test**:

```bash
# Dry run for a Deno repo (shows what would be deleted without running it)
yog reborn deno --dry-run

# Dry run with --deep flag (includes lockfiles)
yog reborn bun --deep --dry-run

# Run clean but skip the package manager install step
yog reborn --no-install
```
**Expected Behavior**:
- It should identify the dummy artifacts and list them for deletion.
- Auto-detects the runtime based on `deno.json` or `package.json` (defaults to Bun for Package.json).

### 5. Testing `yog init`

This command scaffolds a new project from your personal templates.

**Setup**:
1. Create a dummy template at `~/.config/yog/templates/demo/` and add some files inside.

**Commands to test**:

```bash
# Scaffold replacing files to current directory
yog init demo
```

### 6. Testing `yog hack`

Exposes your internal security utility wrapper.

```bash
yog hack analyze
```
**Expected Behavior**:
- Triggers the wrapper and prints a styled execution box using your `hack_go` logic.

---

*Yog creates the world. Azathoth understands it. You do everything.*
