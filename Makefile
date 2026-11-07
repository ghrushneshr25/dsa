# Optional local targets. readme.md is committed by GitHub Actions (.github/workflows/regenerate-docs.yml).
.PHONY: test hooks readme verify-readme

# Generator lives in the sibling dsa-pavilion repo (docgen + Docusaurus site).
DOCGEN_DIR ?= ../dsa-pavilion/docgen

test:
	go test ./...

hooks:
	git config core.hooksPath scripts/git-hooks
	@echo "core.hooksPath=scripts/git-hooks — pre-commit runs go test ./... only."

readme:
	@test -f "$(DOCGEN_DIR)/main.go" || (echo "missing docgen at $(DOCGEN_DIR) — clone dsa-pavilion next to this repo, or set DOCGEN_DIR" >&2; exit 1)
	go -C "$(DOCGEN_DIR)" run . -readme-only -code "$(CURDIR)" -readme "$(CURDIR)/readme.md"
	@echo "readme.md written (commit yourself, or push and let CI update it)"

verify-readme: readme
	git diff --exit-code readme.md || (echo >&2 "readme.md differs from docgen output"; exit 1)
	@echo "readme.md matches docgen"
