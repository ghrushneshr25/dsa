#!/usr/bin/env bash
# Regenerates readme.md via sibling docgen and amends the commit just created so readme
# is part of the same commit (no second manual git add needed).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

DOCGEN_DIR="${DOCGEN_DIR:-$ROOT/../docgen}"
if [ ! -f "$DOCGEN_DIR/go.mod" ] || [ ! -f "$DOCGEN_DIR/main.go" ]; then
  exit 0
fi

echo "post-commit: regenerate readme.md (docgen -readme-only)"
(cd "$DOCGEN_DIR" && go run . -readme-only -code "$ROOT" -readme "$ROOT/readme.md")

# Compare working tree + index to HEAD (covers modified or previously-untracked readme)
if [ -z "$(git status --porcelain -- readme.md)" ]; then
  echo "post-commit: readme.md unchanged — no amend"
  exit 0
fi

echo "post-commit: amend commit to include readme.md"
git add readme.md
git commit --amend --no-edit --no-verify

echo "post-commit: ok (commit amended with readme.md)"
