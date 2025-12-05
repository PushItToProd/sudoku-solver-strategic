#!/usr/bin/env bash

SCRIPT_DIR="$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)"
ROOT_DIR="$SCRIPT_DIR/.."
PUZZLES_DIR="$SCRIPT_DIR/puzzles"

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

solve_sudoku() {
  run_cmd cmd/solve_sudoku "$@"
}

read_puzzle() {
  local puzzle_path="$PUZZLES_DIR"/"$1"
  if [[ ! -f "$puzzle_path" ]]; then
    test::fatal "Unable to get puzzle with id '$1'"
  fi
  tr -d ' \n' <"$puzzle_path"
}

main() {
  output="$(solve_sudoku 2>&1)"
  if [[ "$output" != *'sudoku'* ]]; then
    test::fail "Expected 'sudoku' in cmd/solve_sudoku output when no inputs given"
  fi

  output="$(solve_sudoku 123 2>&1)"
  exit_code=$?
  if (( exit_code == 0 )); then
    test::fail "Expected nonzero exit code when an invalid puzzle is given"
  fi
  if [[ "$output" != *"invalid puzzle"* ]]; then
    test::fail "Expected output to contain 'invalid puzzle'"
  fi

  puzzle_1_solved="$(read_puzzle 1.solved)"
  output="$(solve_sudoku "$puzzle_1_solved")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 1) Expected successful exit code (0) when a valid and solved puzzle is given"
  fi
  if [[ "$output" != *"already solved"* ]]; then
    test::fail "(puzzle 1) Expected output to contain 'already solved'"
  fi

  puzzle_1="$(read_puzzle 1)"
  output="$(solve_sudoku "$puzzle_1")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 1) Expected successful exit code (0) when a valid and solvable puzzle is given"
  fi
  if [[ "$output" == *"already solved"* ]]; then
    test::fail "(puzzle 1) Expected output to not contain 'already solved'"
  fi

  if [[ "$output" != *"$puzzle_1_solved"* ]]; then
    test::fail "(puzzle 1) Expected output to contain the puzzle solution with zeroes replaced with solved digits"
  fi

  puzzle_2_solved="$(read_puzzle 2.solved)"
  output="$(solve_sudoku "$puzzle_2_solved")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 2) Expected successful exit code (0) when a valid and solved puzzle is given"
  fi
  if [[ "$output" != *"already solved"* ]]; then
    test::fail "(puzzle 2) Expected output to contain 'already solved'"
  fi

  puzzle_2="$(read_puzzle 2)"
  output="$(solve_sudoku "$puzzle_2")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 2) Expected successful exit code (0) when a valid and solvable puzzle is given"
  fi
  if [[ "$output" == *"already solved"* ]]; then
    test::fail "(puzzle 2) Expected output to not contain 'already solved'"
  fi

  if [[ "$output" != *"$puzzle_2_solved"* ]]; then
    test::fail "(puzzle 2) Expected output to contain the puzzle solution with zeroes replaced with solved digits"
  fi

  test::check_result
}
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  main "$@"
fi

