#!/bin/bash

# scripts/reborn_bun.sh
set -e

DEEP=$1
DRY_RUN=$2
NO_INSTALL=$3

echo ">> Purging Bun build artifacts..."

TARGETS=("node_modules" ".svelte-kit" ".vite" "dist")

if [ "$DEEP" = "true" ]; then
    TARGETS+=("bun.lockb" "bun.lock")
fi

for TARGET in "${TARGETS[@]}"; do
    if [ -e "$TARGET" ]; then
        if [ "$DRY_RUN" = "true" ]; then
            echo "[DRY-RUN] Would remove: $TARGET"
        else
            echo "Removing: $TARGET"
            rm -rf "$TARGET"
        fi
    fi
done

if [ "$NO_INSTALL" = "true" ]; then
    echo ">> Skipping installation as requested."
    exit 0
fi

if [ "$DRY_RUN" = "true" ]; then
    echo "[DRY-RUN] Would run: bun install"
else
    echo ">> Reinstalling dependencies..."
    bun install
fi

echo ">> Bun reborn complete."
