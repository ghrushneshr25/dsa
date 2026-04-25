#!/usr/bin/env bash
# Run once from the dsa repo root.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"
git config --unset-all core.hooksPath 2>/dev/null || true
cp "$ROOT/scripts/git-hooks/pre-commit" "$ROOT/.git/hooks/pre-commit"
chmod +x "$ROOT/.git/hooks/pre-commit"
echo "Installed .git/hooks/pre-commit (runs scripts/pre-commit.sh)."
