# Optional local targets. readme.md is committed by GitHub Actions (.github/workflows/readme.yml).
.PHONY: test hooks readme verify-readme

test:
	go test ./...

hooks:
	git config core.hooksPath scripts/git-hooks
	@echo "core.hooksPath=scripts/git-hooks — pre-commit runs go test ./... only."

readme:
	@test -f ../docgen/main.go || (echo "missing ../docgen (clone docgen next to this repo, or set DOCGEN_DIR)" >&2; exit 1)
	cd ../docgen && go run . -readme-only -code "$(CURDIR)" -readme "$(CURDIR)/readme.md"
	@echo "readme.md written (commit yourself, or push and let CI update it)"

verify-readme: readme
	git diff --exit-code readme.md || (echo >&2 "readme.md differs from docgen output"; exit 1)
	@echo "readme.md matches docgen"
