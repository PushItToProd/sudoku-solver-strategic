#!/usr/bin/env bash

SCRIPT_DIR="$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)"
ROOT_DIR="$SCRIPT_DIR/.."

failed=0

test::fail() {
  (( failed++ ))
  echo "test failed:" "$@" >&2
}

test::fatal() {
  (( failed++ ))
  echo "fatal:" "$@" >&2
  exit 1
}

test::check_result() {
  if (( failed > 0 )); then
    echo "${failed} tests failed" >&2
    return 1
  else
    echo "All tests passed!"
  fi
}

# _capture_stdout() {
#   __captured_stdout="$(cat)"
# }

# assert::stdout() {
#   local expected="${1?}"
#   _capture_stdout
#   if ! grep -q "$expected" <<<"$__captured_stdout"; then
#     local msg="$2"
#     if [[ ! "$msg" ]]; then
#       msg="expected output to contain "
#     fi
#     test::fail ""
#   fi
# }

run_cmd() {
  local cmd="$1"
  shift
  go run "$ROOT_DIR"/"$cmd" "$@"
}

main() {
  if ! run_cmd cmd/solve_sudoku | grep -qi 'sudoku'; then
    test::fail "Expected 'sudoku' in cmd/solve_sudoku output"
  fi

  test::check_result
}
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  main "$@"
fi

