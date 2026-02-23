import subprocess
import sys
import platform
from pathlib import Path
from importlib.resources import files


def _get_binary() -> Path:
    system = platform.system().lower()

    match platform.machine().lower():
        case "x86_64" | "amd64":
            arch = "amd64"
        case "aarch64" | "arm64":
            arch = "arm64"
        case m:
            raise RuntimeError(f"Unsupported architecture: {m}")

    match (system, arch):
        case ("windows", a):
            name = f"yog-windows-{a}.exe"
        case ("darwin", a):
            name = f"yog-darwin-{a}"
        case ("linux", a):
            name = f"yog-linux-{a}"
        case _:
            raise RuntimeError(f"Unsupported platform: {system}/{arch}")

    try:
        pkg_bin = Path(str(files("yog_sothoth").joinpath(f"bin/{name}")))
        if pkg_bin.exists():
            return pkg_bin
    except Exception:
        pass

    fallback = Path.home() / "go" / "bin" / "yog"
    if fallback.exists():
        return fallback

    raise FileNotFoundError(
        f"yog binary not found for {system}/{arch}.\n"
        f"Expected bundled binary: {name}\n"
        f"For local dev run: just build\n"
        f"For a full package build run: just build-py"
    )


def main() -> None:
    result = subprocess.run([str(_get_binary()), *sys.argv[1:]])
    sys.exit(result.returncode)
