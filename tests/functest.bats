load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'
puzzle_1_solved=$(bash tests/read_puzzle.sh 1.solved)
puzzle_1=$(bash tests/read_puzzle.sh 1)
puzzle_2_invalid=$(bash tests/read_puzzle.sh 2.invalid)
puzzle_2_solved=$(bash tests/read_puzzle.sh 2.solved)
puzzle_2=$(bash tests/read_puzzle.sh 2)

@test "running solve_sudoku with no arguments prints 'sudoku' and returns nonzero" {
  run go run ./cmd/solve_sudoku
  assert_failure
  assert_output --partial 'sudoku'
}

@test "'solve_sudoku 123' returns an error" {
  run go run ./cmd/solve_sudoku 123
  assert_failure
  assert_output --partial 'too short'
}

@test "report if a puzzle is already solved" {
  run go run ./cmd/solve_sudoku "$puzzle_1_solved"
  assert_success
  assert_output --partial 'already solved'
}

@test "an unfinished and solvable puzzle is solved" {
  run go run ./cmd/solve_sudoku "$puzzle_1"
  assert_success
  assert_output --partial "$puzzle_1_solved"
}

@test "invalid puzzles are reported" {
  run go run ./cmd/solve_sudoku "$puzzle_2_invalid"
  assert_failure
  assert_output --partial 'invalid puzzle'
}

@test "puzzle 2 is correctly solved" {
  run go run ./cmd/solve_sudoku "$puzzle_2"
  assert_success
  assert_output --partial "$puzzle_2_solved"
}
