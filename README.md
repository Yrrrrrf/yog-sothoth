# Yog-Sothoth

> _The Key and the Gate_

A personal infrastructure CLI tool written in Go. Not something you work
_inside_ — it's the ritual you run **at the threshold** of a project. It
eliminates friction at project lifecycle boundaries: starting clean, staying
clean, validating everything.

Pairs with **[Azathoth](https://github.com/Yrrrrrf/azathoth)** (AI/MCP layer) as
a personal dev acceleration system.

> _Yog creates the world. Azathoth understands it. You do everything._

---

## Commands

| Command                | Description                                                        |
| ---------------------- | ------------------------------------------------------------------ |
| `yog doctor`           | Health check — runtimes, `.env` completeness, git, README, LICENSE |
| `yog env check`        | Validates nothing is missing vs `.env.example`                     |
| `yog env diff`         | Surfaces drift — vars in `.env` but not in the template            |
| `yog env sync`         | Interactively fills in missing variables                           |
| `yog env load`         | Walks up the directory tree to find and load `.env`                |
| `yog reborn [runtime]` | Deep-cleans build artifacts and reinstalls from scratch            |

`reborn` flags: `--deep` (also nuke lockfiles), `--full` (recursive workspace
clean), `--dry-run`, `--no-install`

> **Node.js is never supported.** Deno and Bun only.

---

## Install

### Go binary

```bash
go build -o ~/go/bin/yog ./src/yog_sothoth
```

Ensure `~/go/bin` is in your `$PATH`. On NixOS, add it via `home.sessionPath` in
home-manager, or use the uv wrapper below.

### uv wrapper (NixOS / read-only systems)

```bash
uv tool install --editable .
```

This installs a `yog` shim via the Python wrapper that proxies to the compiled
Go binary — no PATH surgery required.

---

## Architecture

Strict 3-layer separation — non-negotiable:

```
src/yog_sothoth/
├── main.go          ← calls cmd.Execute(), nothing else
├── cmd/             ← CLI layer: Cobra only, zero logic
├── internal/
│   ├── deps/        ← reborn logic
│   ├── env/         ← .env management logic
│   └── health/      ← doctor check logic
└── pkg/             ← shared config types + Lip Gloss UI styles
```

- `cmd/` defines flags and delegates. Never contains business logic.
- `internal/` does all real work. Never knows about Cobra.
- `pkg/` holds only what's genuinely shared across layers.

---

## Tech Stack

- **Language:** Go — [Cobra](https://github.com/spf13/cobra) ·
  [Viper](https://github.com/spf13/viper) ·
  [Lip Gloss](https://github.com/charmbracelet/lipgloss) ·
  [godotenv](https://github.com/joho/godotenv)
- **Python shim:** uv workspace · `uv tool install --editable .`

---

## License

This project is licensed under the terms of the [MIT license](LICENSE).
