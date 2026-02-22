# Yog-Sothoth — Vision & Context
> *The Key and the Gate*

## What it is

Yog-Sothoth is a personal **infrastructure CLI tool** written in Go. It is not a tool you work *inside* — it's the ritual you run *at the threshold*, before a project begins. It eliminates all friction at project lifecycle boundaries: starting, onboarding, recovering, validating.

It pairs with a companion tool called **Azathoth** (an AI/MCP layer). Together they form a personal dev acceleration system:
- **Yog creates the world. Azathoth understands it. You do everything.**

## Current version: `0.0.2`

## Tech stack

- **Language:** Go 1.25+
- **CLI framework:** Cobra (`github.com/spf13/cobra`)
- **Config:** Viper (`github.com/spf13/viper`)
- **UI/TUI:** Charm ecosystem — Lip Gloss, Bubbletea, Glamour
- **Env parsing:** `joho/godotenv`
- **Bash scripts:** embedded directly into the binary via `//go:embed`

## Architecture (3-layer separation — hard rule)

```
src/yog_sothoth/
├── main.go              ← calls cmd.Execute() and nothing else
├── cmd/                 ← CLI layer: Cobra lives here ONLY
├── internal/            ← Logic layer: no Cobra, no CLI concerns
│   ├── deps/            ← reborn logic
│   ├── env/             ← .env management logic
│   ├── hack/            ← thin wrapper over hack_go library
│   ├── health/          ← doctor check logic
│   └── scaffold/        ← init/template logic
├── pkg/                 ← shared: config types, Lip Gloss styles
└── scripts/             ← bash scripts (embedded into binary)
```

**The rules:**
- `cmd/` defines flags and calls `internal/`. Never contains logic.
- `internal/` does all real work. Never knows about Cobra.
- `pkg/` holds only what's genuinely shared across both layers.
- `scripts/` are standalone bash, embedded via `//go:embed`, never imported.

## Commands

| Command | What it does |
|---|---|
| `yog init [template]` | Scaffolds a project from `~/.config/yog/templates/`, then auto-runs reborn |
| `yog reborn [runtime]` | Deep-cleans build artifacts and reinstalls deps. Flags: `--deep`, `--dry-run`, `--no-install` |
| `yog env check/diff/sync/load` | .env validation, drift detection, interactive sync |
| `yog doctor` | Full health check: runtimes, tools, .env, git, README, LICENSE |
| `yog hack [tool]` | Exposes `hack_go` library — cybersecurity learning tools, authorized use only |

## Key constraints (non-negotiable)

- **Node.js is never supported.** Deno and Bun only.
- **Every command must solve a friction point that occurs 3+ times/week.** No feature bloat.
- **`cmd/` never contains business logic.** Ever.
- All terminal styling goes through `pkg/ui/` — never inline in `cmd/`.

## Local install (the "editable" workflow)

```bash
go build -o ~/go/bin/yog ./src/yog_sothoth
# or
go install ./src/yog_sothoth
```

Requires `~/go/bin` in `$PATH`. Rebuild after changes — that's it.

## Companion: hack_go

`src/hack_go` is a separate Go library (linked via `go.work`) that contains cybersecurity learning tools. `internal/hack/` is the thin adapter that exposes it through the CLI. It's intentionally minimal today — it's the long-term growth area of the project.

## What's done vs. what's next

**Done:** Full architecture, `doctor` ✓, `reborn` ✓, `env` ✓, `hack` (stub) ✓, binary embeds, Charm UI, local install workflow.

**Next:** Template variable substitution in `init`, real tool implementations inside `hack_go`.