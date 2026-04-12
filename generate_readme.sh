#!/bin/bash

ROOT_DIR="./codebase"
OUTPUT="readme.md"

echo "Generating README..."

# ----------------------------
# FORMAT CATEGORY TITLE
# ----------------------------
format_header() {
  local folder="$1"
  folder=$(basename "$folder")
  echo "$folder" | tr '-' ' ' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)}1'
}

# ----------------------------
# EXTRACT METADATA
# ----------------------------
extract_meta() {
  local key="$1"
  local file="$2"

  grep "^// @$key:" "$file" | sed "s|// @$key:||" | xargs
}

# ----------------------------
# EXTRACT DESCRIPTION SECTION
# ----------------------------
extract_description() {
  local file="$1"

  awk '
    /@section: Description/ {flag=1; next}
    /@section:/ {if(flag) exit}
    flag {
      gsub(/^[ \t\*]+/, "")
      printf "%s ", $0
    }
  ' "$file"
}

# ----------------------------
# GENERATE README
# ----------------------------
{
  echo "# 📚 DSA Index"
  echo ""

  shopt -s nullglob

  for dir in "$ROOT_DIR"/*/; do
    [[ "$dir" == *".git"* ]] && continue

    files=("$dir"/*.go)
    [[ ${#files[@]} -eq 0 ]] && continue

    section_name=$(format_header "$dir")

    echo "## 📂 $section_name"
    echo ""
    echo "| Problem | Difficulty | Tags | Description | Code | Tests |"
    echo "|--------|------------|------|-------------|------|-------|"

    TEST_DIR="${dir%/}/test"

    # SORT FILES (problem1, problem2...)
    IFS=$'\n' files=($(ls "$dir"/*.go | sort -V))
    unset IFS

    for file in "${files[@]}"; do
      [[ "$file" == *_test.go ]] && continue

      filename=$(basename "$file")
      name="${filename%.go}"

      # ----------------------------
      # METADATA
      # ----------------------------
      problem_name=$(extract_meta "problem" "$file")
      difficulty=$(extract_meta "difficulty" "$file")
      tags=$(extract_meta "tags" "$file")

      [[ -z "$problem_name" ]] && problem_name="$name"

      # ----------------------------
      # DESCRIPTION
      # ----------------------------
      description=$(extract_description "$file")
      [[ -z "$description" ]] && description="—"

      # limit length (clean UI)
      description=$(echo "$description" | cut -c1-120)

      # ----------------------------
      # LINKS
      # ----------------------------
      code_link="[Link](${file})"

      test_file="$TEST_DIR/${name}_test.go"

      if [[ -f "$test_file" ]]; then
        test_link="[Link](${test_file})"
      else
        test_link="—"
      fi

      # ----------------------------
      # ROW
      # ----------------------------
      echo "| $problem_name | $difficulty | $tags | $description | $code_link | $test_link |"

    done

    echo ""
  done

} > "$OUTPUT"

echo "README.md generated successfully!"