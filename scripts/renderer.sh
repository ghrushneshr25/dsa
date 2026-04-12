#!/opt/homebrew/bin/bash

render_index_header() {
  local file=$1
  local title=$2
  local slug=$3

  cat <<EOF > "$file"
---
title: $title
slug: /$slug
pagination_prev: null
pagination_next: null
---

# 📂 $title

| Problem | Difficulty | Tags |
|--------|------------|------|
EOF
}

render_index_row() {
  local file=$1
  local title=$2
  local link=$3
  local difficulty=$4
  local tags=$5

  echo "| [$title]($link) | $difficulty | $tags |" >> "$file"
}

render_doc() {
  local file=$1
  local title=$2
  local difficulty=$3
  local tags=$4
  local time=$5
  local space=$6
  local sections=$7
  local code=$8
  local tests=$9
  local source=${10}

  cat <<EOF > "$file"
---
title: $title
---

**Difficulty:** $difficulty  
**Tags:** $tags  

**Time:** $time  
**Space:** $space  

$sections

## Code (Go)

\`\`\`go
$code
\`\`\`

$tests

[View Source]($source)

EOF
}