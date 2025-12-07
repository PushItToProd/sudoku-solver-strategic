#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PUZZLES_DIR="$SCRIPT_DIR/puzzles"

# Remove all whitespace and newlines
squash() {
  tr -d ' \n'
}

strip_comments() {
  grep -v '^#'
}

read_puzzle() {
  local puzzle_path="$PUZZLES_DIR"/"$1"
  if [[ ! -f "$puzzle_path" ]]; then
    test::fatal "Unable to get puzzle with id '$1'"
  fi
  strip_comments <"$puzzle_path" | squash
}

main() {
  local puzzle="${1?must provide a puzzle ID}"
  read_puzzle "$puzzle"
  if [[ -t 1 ]]; then
    echo
  fi
}
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  main "$@"
fi
