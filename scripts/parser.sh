#!/opt/homebrew/bin/bash

parse_metadata() {
  local file=$1
  declare -n meta_ref=$2

  while IFS= read -r line; do
    if [[ $line =~ ^//\ @([a-zA-Z0-9_-]+):\ (.*)$ ]]; then
      meta_ref["${BASH_REMATCH[1]}"]="${BASH_REMATCH[2]}"
    fi
  done < "$file"
}

parse_sections() {
  local file=$1

  raw=$(sed -n '/\/\*/,/\*\//p' "$file" | sed '1d;$d')

  formatted=""
  current_title=""
  current_content=""

  while IFS= read -r line; do
    if [[ $line =~ ^@section:\ (.*) ]]; then
      if [ -n "$current_title" ]; then
        if [[ "$current_title" == "Algorithm" ]]; then
          formatted+=$'\n## Algorithm\n\n```text\n'"$current_content"$'\n```\n'
        else
          formatted+=$'\n## '"$current_title"$'\n\n'"$current_content"$'\n'
        fi
      fi

      current_title="${BASH_REMATCH[1]}"
      current_content=""
    else
      current_content+=$(echo "$line" | sed 's/^ *//')$'\n'
    fi
  done <<< "$raw"

  echo "$formatted"
}

parse_tests() {
  local file=$1

  output=""
  inside=0
  brace_count=0
  title=""
  desc=""
  code=""

  while IFS= read -r line; do

    # detect start
    if [[ $line =~ t\.Run\(\"([^\"]+)\" ]]; then
      inside=1
      title="${BASH_REMATCH[1]}"
      desc=""
      code="$line"$'\n'
      brace_count=1
      continue
    fi

    if [ $inside -eq 1 ]; then

      # capture description
      if [[ $line =~ @desc:\ (.*) ]]; then
        desc="${BASH_REMATCH[1]}"
        continue
      fi

      # count braces
      open=$(grep -o '{' <<< "$line" | wc -l | xargs)
      close=$(grep -o '}' <<< "$line" | wc -l | xargs)

      brace_count=$((brace_count + open - close))

      code+="$line"$'\n'

      # end of block
      if [ $brace_count -eq 0 ]; then
        inside=0

        output+=$'\n### '"$title"$'\n'

        if [ -n "$desc" ]; then
          output+=$'\n'"$desc"$'\n'
        fi

        output+=$'\n```go\n'"$code"$'\n```\n'
      fi
    fi

  done < "$file"

  echo "$output"
}