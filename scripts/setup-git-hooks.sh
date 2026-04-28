#!/usr/bin/env bash
# Run once from the dsa repo root: versioned hooks via core.hooksPath (no copy into .git/hooks).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"
git config --unset-all core.hooksPath 2>/dev/null || true
git config core.hooksPath scripts/git-hooks
# Legacy install: remove copied hook so hooksPath is the only source
rm -f "$ROOT/.git/hooks/pre-commit"
chmod +x "$ROOT/scripts/git-hooks/pre-commit" 2>/dev/null || true
echo "Installed core.hooksPath=scripts/git-hooks (runs scripts/git-hooks/pre-commit on commit)."
