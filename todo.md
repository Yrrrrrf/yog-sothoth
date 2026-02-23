# Yog-Sothoth — Next Improvements

---

## Fix 1 — `yog env load` shell integration

> _The CLI can't mutate its parent shell. Change the contract: emit, don't load._

The current implementation calls `os.Setenv()` inside the Go process, which only
affects the child process and is lost the moment the command exits. The fix is to
make `env load` print `export` statements to stdout and let the shell evaluate them.

### Changes needed

- [ ] Rewrite `envLoadCmd` in `cmd/env.go` to print to stdout instead of setting env vars:
  ```go
  for k, v := range vars {
      fmt.Printf("export %s=%q\n", k, v)
  }
  ```
- [ ] Remove any `os.Setenv()` calls from the load path
- [ ] Remove the `--show-values` flag (now redundant — stdout is the output)
- [ ] Update `README.md` to document the new usage pattern:
  ```bash
  # Inline eval
  eval "$(yog env load)"

  # Or add a shell alias (recommended)
  alias yog-load='eval "$(yog env load)"'
  ```
- [ ] Add a note in `yog env load --help` explaining the eval pattern and why
  direct loading is not possible from a subprocess

### References

Same pattern used by `direnv`, `rbenv`, `nvm`, and `ssh-agent`.

---

## Fix 2 — Stop tracking compiled binaries in git

> _Binaries are build artifacts. They belong in `.gitignore`, not in history._

Every `just build` produces new `yog-*` binaries in `src/yog_sothoth/bin/`. Since
they are re-compilable at any time, there is no reason to track them — they bloat
the repo and the package with every rebuild.

### Changes needed

- [ ] Add the following to `.gitignore`:
  ```
  src/yog_sothoth/bin/yog-*
  ```
- [ ] Add a `src/yog_sothoth/bin/.gitkeep` so the `bin/` directory itself stays
  tracked (git does not track empty directories)
- [ ] Untrack any binaries already committed to history:
  ```bash
  git rm --cached src/yog_sothoth/bin/yog-*
  git commit -m "chore: untrack compiled binaries from git history"
  ```
- [ ] Verify `pyproject.toml` still packages them correctly at wheel build time —
  the `include = ["src/yog_sothoth/bin/*"]` glob is fine because `just build`
  will have produced the binaries on disk before `uv build` runs

### Notes

- `just clean-bin` already exists — use it after confirming gitignore is in place
- CI (Phase 5 from old todo) must run `just build` before `uv build` so the
  binaries exist at publish time even though they are not in the repo