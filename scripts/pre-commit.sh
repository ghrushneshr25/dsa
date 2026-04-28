#!/usr/bin/env bash
# Runs go test ./... before each commit (readme.md is regenerated in CI: .github/workflows/readme.yml).
#
# One-time: ./scripts/setup-git-hooks.sh   or   make hooks
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "pre-commit: go test ./... (module root)"
go test ./...

echo "pre-commit: ok"
