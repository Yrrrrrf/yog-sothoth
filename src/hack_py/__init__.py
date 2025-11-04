# ...existing code...
from __future__ import annotations

from typing import Optional

try:
    from importlib import metadata as importlib_metadata  # py3.8+
except Exception:
    import importlib_metadata  # type: ignore

from rich.console import Console

PROJECT_NAME = "hack_py"


def get_version(package: Optional[str] = None) -> str:
    pkg = package or PROJECT_NAME
    try:
        return importlib_metadata.version(pkg)
    except Exception:
        return "0.0.0"


def print_banner(name: Optional[str] = None, version: Optional[str] = None, console: Optional[Console] = None) -> None:
    console = console or Console()
    name = name or PROJECT_NAME
    version = version or get_version(name)
    console.print(f"[green]{name}[/green] [blue italic]{version}[/blue italic]")
    console.print("some print str...")


if __name__ == "__main__":
    print_banner()
# ...existing code...

def main():
    print("Hello from hack_py!")
    print("Hello from hack_py!")
    print("Hello from hack_py!")
    print("Hello from hack_py!")
    print("Hello from hack_py!")

if __name__ == "__main__":
    main()
