#!/usr/bin/env bash
# Runs tests, then regenerates readme.md when docgen is available and stages it
# for the same commit (run from repo root; typical layout: dsa-doc/dsa + dsa-doc/docgen).
#
# One-time: ./scripts/setup-git-hooks.sh   or   make hooks
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "pre-commit: go test ./... (module root)"
go test ./...

DOCGEN_DIR="${DOCGEN_DIR:-}"
if [ -z "$DOCGEN_DIR" ]; then
  for candidate in "$ROOT/../docgen" "$ROOT/../../docgen"; do
    if [ -f "$candidate/go.mod" ] && [ -f "$candidate/main.go" ]; then
      DOCGEN_DIR="$candidate"
      break
    fi
  done
fi

if [ -n "$DOCGEN_DIR" ] && [ -f "$DOCGEN_DIR/go.mod" ] && [ -f "$DOCGEN_DIR/main.go" ]; then
  echo "pre-commit: regenerate readme.md (docgen -readme-only, DOCGEN_DIR=$DOCGEN_DIR)"
  (cd "$DOCGEN_DIR" && go run . -readme-only -code "$ROOT" -readme "$ROOT/readme.md")
  git add readme.md
else
  echo "pre-commit: skip readme regen (no docgen; tried ../docgen and ../../docgen; set DOCGEN_DIR). CI still checks readme via .github/workflows/test.yml"
fi

echo "pre-commit: ok"
