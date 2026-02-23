# Embed `yog` binary into the Python package

> _Make `uv tool install yog-sothoth` the only install step anyone ever needs_

The goal: ship the compiled Go binary **inside** the PyPI package so users get a
fully functional `yog` CLI with zero Go toolchain requirements. Same pattern
used by `ruff`, `pyright`, and `esbuild`.

---

## Phase 1 — Restructure the package

- [ ] Create `src/yog_sothoth/bin/` directory
- [ ] Add a `.gitkeep` so the folder is tracked but binaries are gitignored
- [ ] Update `.gitignore` to exclude `src/yog_sothoth/bin/yog-*`
- [ ] Add `src/yog_sothoth/bin/__init__.py` (empty, makes it a proper
      subpackage)
- [ ] Wire up package data in `pyproject.toml`:
  ```toml
  [tool.setuptools.package-data]
  "yog_sothoth.bin" = ["*"]
  ```

---

## Phase 2 — Update `cli.py`

- [ ] Replace hardcoded `Path.home() / "go" / "bin" / "yog"` with dynamic
      resolution
- [ ] Implement `get_binary()` using `platform.system()` + `platform.machine()`
- [ ] Use `importlib.resources.files()` to resolve the binary path inside the
      package
- [ ] Handle the `.exe` extension on Windows
- [ ] Handle `aarch64` → `arm64` machine name normalization
- [ ] Add a clear error if the binary for the current platform is missing (don't
      silently fail)
- [ ] Keep the old `Path.home()` fallback for local dev (editable installs
      before a build)

```
# Target matrix to support:
yog-linux-amd64
yog-linux-arm64
yog-darwin-amd64
yog-darwin-arm64
yog-windows-amd64.exe
```

---

## Phase 3 — Build script

- [ ] Write `build.sh` at the repo root that cross-compiles all targets
- [ ] Iterate over `GOOS` × `GOARCH` combinations
- [ ] Output binaries directly into `src/yog_sothoth/bin/`
- [ ] Set executable permissions (`chmod +x`) on non-Windows binaries
- [ ] Print a summary of file sizes after build (sanity check — expect ~8-12MB
      each)
- [ ] Make the script fail fast on any compilation error (`set -e`)

---

## Phase 4 — Publish pipeline

- [ ] Add `build.sh` as a prerequisite step before `uv build`
- [ ] Verify the built wheel actually contains the binaries
      (`unzip -l dist/*.whl | grep bin/`)
- [ ] Test install from the wheel locally:
      `uv tool install dist/yog_sothoth-*.whl`
- [ ] Confirm `yog --help` works after wheel install with no Go on PATH
- [ ] Publish: `uv publish`

---

## Phase 5 — CI (optional but nice)

- [ ] Add a GitHub Actions workflow that:
  - triggers on version tag push (`v*`)
  - runs `build.sh` to cross-compile all targets
  - runs `uv build`
  - runs `uv publish` with PyPI token from secrets
- [ ] Cache the Go build artifacts between runs

---

## Open questions

- **Windows arm64?** Go supports it. Probably skip for now unless there's a real
  use case.
- **Binary size:** Strip debug info with `-ldflags="-s -w"` to cut ~30% off each
  binary. Worth it.
- **Editable installs:** `uv tool install --editable .` won't have the binaries
  pre-built — the fallback to `~/go/bin/yog` covers this, but document it
  clearly.
- **Version sync:** Should the Python package version and the Go binary version
  be kept in lockstep? Yes. Automate it or it will drift.

---

## Reference implementations to study

- `ruff` — platform wheels approach (more complex, not needed here)
- `pyright` — single package with bundled binary, very close to what we want
- `esbuild-python` — simple subprocess wrapper with embedded binary, basically
  this exact pattern

---

_When this is done: `uv tool install yog-sothoth` is the only command anyone
needs. No Go. No PATH. No friction._
