#!/opt/homebrew/bin/bash

format_title() {
  echo "$1" | sed -E 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)}1'
}

slugify() {
  echo "$1" | \
    tr '[:upper:]' '[:lower:]' | \
    sed -E 's/[^a-z0-9]+/-/g' | \
    sed -E 's/^-+|-+$//g'
}