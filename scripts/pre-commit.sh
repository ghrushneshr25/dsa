#!/usr/bin/env bash
# Runs tests before each commit. Readme refresh happens in post-commit (amend) so it
# lands in the same commit even if the IDE only staged some files.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "pre-commit: go test ./... (module root)"
go test ./...

echo "pre-commit: ok"
