#!/opt/homebrew/bin/bash

init_sidebar() {
  local file=$1

  echo "// @ts-check" > "$file"
  echo "/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */" >> "$file"
  echo "const sidebars = { tutorialSidebar: [" >> "$file"
}

start_category() {
  local file=$1
  local label=$2
  local id=$3

  cat <<EOF >> "$file"
{
  type: 'category',
  label: '$label',
  link: { type: 'doc', id: '$id/index' },
  items: [
EOF
}

add_sidebar_item() {
  local file=$1
  local item=$2

  echo "    '$item'," >> "$file"
}

end_category() {
  local file=$1

  echo "  ]" >> "$file"
  echo "}," >> "$file"
}

close_sidebar() {
  local file=$1

  echo "] };" >> "$file"
  echo "export default sidebars;" >> "$file"
}