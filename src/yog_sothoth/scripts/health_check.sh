#!/bin/bash

# scripts/health_check.sh
# This script can be run standalone to perform underlying system checks outside of the Cobra UX.
# It checks tool versions. The Cobra CLI relies primarily on native Go checks for a richer UX.

echo "--- Yog-Sothoth Basic Health Check ---"
if command -v deno >/dev/null 2>&1; then
    echo "Deno: $(deno --version | head -n 1)"
else
    echo "Deno: Not installed"
fi

if command -v bun >/dev/null 2>&1; then
    echo "Bun: $(bun --version)"
else
    echo "Bun: Not installed"
fi

if command -v git >/dev/null 2>&1; then
    echo "Git: $(git --version)"
else
    echo "Git: Not installed"
fi
