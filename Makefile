# Run from the dsa repo root. Readme regen expects docgen as a sibling directory
# (…/dsa-doc/dsa and …/dsa-doc/docgen), same as scripts/pre-commit.sh.
.PHONY: test hooks readme verify-readme

test:
	go test ./...

hooks:
	git config core.hooksPath scripts/git-hooks
	@echo "core.hooksPath=scripts/git-hooks — pre-commit runs tests + readme regen when ../docgen exists."

readme:
	@test -f ../docgen/main.go || (echo "missing ../docgen (clone docgen next to this repo, or set DOCGEN_DIR)" >&2; exit 1)
	cd ../docgen && go run . -readme-only -code "$(CURDIR)" -readme "$(CURDIR)/readme.md"
	@echo "readme.md written"

verify-readme: readme
	git diff --exit-code readme.md || (echo >&2 "readme.md differs from docgen output; commit after: make readme"; exit 1)
	@echo "readme.md matches docgen"
