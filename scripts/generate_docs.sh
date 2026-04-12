#!/opt/homebrew/bin/bash
set -e

source scripts/utils.sh
source scripts/sidebar.sh
source scripts/parser.sh
source scripts/renderer.sh

SIDEBAR_FILE="website/sidebars.js"
BASE_CODE_DIR="codebase"
BASE_DOCS_DIR="website/docs"
REPO_URL="https://github.com/ghrushneshr25/dsa"

mkdir -p "$BASE_DOCS_DIR"
rm -rf "$BASE_DOCS_DIR"/*

init_sidebar "$SIDEBAR_FILE"

for category_path in "$BASE_CODE_DIR"/*; do
  [ -d "$category_path" ] || continue

  raw_category=$(basename "$category_path")
  category=$(format_title "$raw_category")

  DOCS_DIR="$BASE_DOCS_DIR/$raw_category"
  mkdir -p "$DOCS_DIR"

  INDEX_FILE="$DOCS_DIR/index.md"

  start_category "$SIDEBAR_FILE" "$category" "$raw_category"
  render_index_header "$INDEX_FILE" "$category" "$raw_category"

  for file in $(ls "$category_path"/*.go | sort -V); do
    [[ "$file" == *_test.go ]] && continue

    filename=$(basename "$file")
    declare -A meta

    parse_metadata "$file" meta

    title="${meta[problem]}"
    slug=$(slugify "$title")

    add_sidebar_item "$SIDEBAR_FILE" "$raw_category/$slug"

    render_index_row "$INDEX_FILE" "$title" "/dsa/$raw_category/$slug" "${meta[difficulty]}" "${meta[tags]}"

    sections=$(parse_sections "$file")
    code=$(sed '/^\/\/ @/d;/\/\*/,/\*\//d' "$file")

# ----------------------------
# TESTS (structured parsing)
# ----------------------------
test_file_root="$category_path/${filename%.go}_test.go"
test_file_nested="$category_path/test/${filename%.go}_test.go"

tests=""

if [ -f "$test_file_root" ]; then
  parsed=$(parse_tests "$test_file_root")

elif [ -f "$test_file_nested" ]; then
  parsed=$(parse_tests "$test_file_nested")

fi

if [ -n "$parsed" ]; then
  tests=$(cat <<EOF

## Tests
$parsed

EOF
)
fi

    render_doc \
      "$DOCS_DIR/$slug.md" \
      "$title" \
      "${meta[difficulty]}" \
      "${meta[tags]}" \
      "${meta[time]}" \
      "${meta[space]}" \
      "$sections" \
      "$code" \
      "$tests" \
      "$REPO_URL/blob/master/codebase/$raw_category/$filename"

  done

  end_category "$SIDEBAR_FILE"

done

close_sidebar "$SIDEBAR_FILE"

echo "✅ Modular doc generator ready"