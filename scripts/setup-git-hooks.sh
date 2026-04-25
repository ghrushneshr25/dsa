#!/usr/bin/env bash
# Run once from the dsa repo root.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"
git config --unset-all core.hooksPath 2>/dev/null || true
cp "$ROOT/scripts/git-hooks/pre-commit" "$ROOT/.git/hooks/pre-commit"
cp "$ROOT/scripts/git-hooks/post-commit" "$ROOT/.git/hooks/post-commit"
chmod +x "$ROOT/.git/hooks/pre-commit" "$ROOT/.git/hooks/post-commit"
echo "Installed .git/hooks/pre-commit and post-commit."
