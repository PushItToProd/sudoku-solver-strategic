#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
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
  run_cmd cmd/solve_sudoku "$@" 2>&1
}

squash_arg() {
  squash_stdin <<<"$1"
}

squash_stdin() {
  tr -d ' \n'
}

read_puzzle() {
  local puzzle_path="$PUZZLES_DIR"/"$1"
  if [[ ! -f "$puzzle_path" ]]; then
    test::fatal "Unable to get puzzle with id '$1'"
  fi
  squash_stdin <"$puzzle_path"
}

main() {
  # Running the program with no inputs emits some kind of help text
  output="$(solve_sudoku)"
  exit_code=$?
  if (( exit_code == 0 )); then
    test::fail "Expected nonzero exit code when no puzzle is given"
    echo "actual output: $output"
  fi
  if [[ "$output" != *'sudoku'* ]]; then
    test::fail "Expected 'sudoku' in cmd/solve_sudoku output when no inputs given"
    echo "actual output: $output"
  fi

  output="$(solve_sudoku 123)"
  exit_code=$?
  if (( exit_code == 0 )); then
    test::fail "Expected nonzero exit code when an invalid puzzle is given"
  fi
  if [[ "$output" != *"too short"* ]]; then
    test::fail "Expected output to contain 'too short'"
    echo "actual output: $output"
  fi

  # puzzle 1
  puzzle_1_solved="$(read_puzzle 1.solved)"
  output="$(solve_sudoku "$puzzle_1_solved")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 1.solved) Expected successful exit code (0) when a valid and solved puzzle is given"
  fi
  if [[ "$output" != *"already solved"* ]]; then
    test::fail "(puzzle 1.solved) Expected output to contain 'already solved'"
    echo "actual output: $output"
  fi

  puzzle_1="$(read_puzzle 1)"
  output="$(solve_sudoku "$puzzle_1")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 1) Expected successful exit code (0) when a valid and solvable puzzle is given"
  fi
  if [[ "$output" == *"already solved"* ]]; then
    test::fail "(puzzle 1) Expected output to not contain 'already solved'"
    echo "actual output: $output"
  fi
  if [[ "$(squash_arg "$output")" != *"$puzzle_1_solved"* ]]; then
    test::fail "(puzzle 1) Expected output to contain the puzzle solution with zeroes replaced with solved digits"
    echo "actual output: $output"
  fi

  # puzzle 2 - invalid
  puzzle2_invalid="$(read_puzzle 2.invalid)"
  output="$(solve_sudoku "$puzzle2_invalid")"
  exit_code=$?
  if (( exit_code == 0 )); then
    test::fail "(puzzle 2.invalid) Expected nonzero exit code when an unsolvable puzzle is given"
  fi
  if [[ "$output" != *"invalid puzzle"* ]]; then
    test::fail "(puzzle 2.invalid) Expected output to contain 'invalid puzzle'"
    echo "actual output: $output"
  fi


  # puzzle 2
  puzzle_2_solved="$(read_puzzle 2.solved)"
  output="$(solve_sudoku "$puzzle_2_solved")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 2.solved) Expected successful exit code (0) when a valid and solved puzzle is given"
  fi
  if [[ "$output" != *"already solved"* ]]; then
    test::fail "(puzzle 2.solved) Expected output to contain 'already solved'"
    echo "actual output: $output"
  fi

  puzzle_2="$(read_puzzle 2)"
  output="$(solve_sudoku "$puzzle_2")"
  exit_code=$?
  if (( exit_code != 0 )); then
    test::fail "(puzzle 2) Expected successful exit code (0) when a valid and solvable puzzle is given"
  fi
  if [[ "$output" == *"already solved"* ]]; then
    test::fail "(puzzle 2) Expected output to not contain 'already solved'"
    echo "actual output: $output"
  fi
  if [[ "$(squash_arg "$output")" != *"$puzzle_2_solved"* ]]; then
    test::fail "(puzzle 2) Expected output to contain the puzzle solution with zeroes replaced with solved digits"
    echo "actual output: $output"
  fi

  test::check_result
}
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  main "$@"
fi

