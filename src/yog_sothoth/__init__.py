import subprocess
import sys
from pathlib import Path

def main() -> None:
    binary = Path.home() / "go" / "bin" / "yog"
    result = subprocess.run([str(binary), *sys.argv[1:]])
    sys.exit(result.returncode)
