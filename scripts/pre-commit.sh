#!/usr/bin/env bash
# Runs tests, then regenerates readme.md when a sibling docgen/ exists and stages it
# for the same commit (run from repo root; typical layout: dsa-doc/dsa + dsa-doc/docgen).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "pre-commit: go test ./... (module root)"
go test ./...

DOCGEN_DIR="${DOCGEN_DIR:-$ROOT/../docgen}"
if [ -f "$DOCGEN_DIR/go.mod" ] && [ -f "$DOCGEN_DIR/main.go" ]; then
  echo "pre-commit: regenerate readme.md (docgen -readme-only)"
  (cd "$DOCGEN_DIR" && go run . -readme-only -code "$ROOT" -readme "$ROOT/readme.md")
  git add readme.md
else
  echo "pre-commit: skip readme regen (no sibling docgen at $DOCGEN_DIR; set DOCGEN_DIR to override)"
fi

echo "pre-commit: ok"
