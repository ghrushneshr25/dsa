#!/usr/bin/env bash
# DSA code repository only: run tests. Doc generation is cross-repo; use docgen
# locally (see docgen/generate.sh) or GitHub Actions on push.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "pre-commit: go test ./... (module root)"
go test ./...

echo "pre-commit: ok"
