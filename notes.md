

---

sudoku states

- valid - is exactly 81 characters that are either valid digits or not yet filled
- complete - all cells are filled in
- correct - has no duplicated or invalid digits (may be incomplete)
- solved - is valid, complete, and correct

---

## 2025-12-06

### Recap

So far I have:

- A set of functional tests written in bash that check basic behavior of the entrypoint `solve_sudoku`
  - `solve_sudoku` takes a sudoku puzzle on the command line as a string of 81 characters
    - it prints "sudoku" if isn't passed an argument
    - it prints "invalid puzzle" if its first argument isn't exactly 81 chars
    - it prints "unsolvable" if `sudoku.New(args[1])` returns an error
    - it prints "already solved" if the given puzzle is already solved
    - it's supposed to solve the puzzle it's given, but so far I've cheated and just have a couple possible inputs hardcoded
- Unit tests for a file `validator.go`
  - don't have a clear interface right now -- I implemented a couple redundant functions
    - `New` creates a `BoardState` or returns an error
    - `IsValidSudoku` just checks if a sudoku is valid
    - `Check` takes a puzzle as a string and returns a `SudokuState` object
  - idea: convert `SudokuState` to a set of error types or to a set of bit flags -- e.g. `ErrInvalidLength`, `ErrInvalidChars`, `ErrIncomplete`, `ErrConflictingDigits`
  - **remember:** the goal of this is to be able to provide useful error messages at the command line

### Main open items, ideas, etc.

- [x] implement solving: use my bruteforce solver implementation to get output working for now
- `solve_sudoku` entrypoint
  - [x] validation and error reporting - `solve_sudoku` should rely on `sudoku.New()` to report error states
  - [x] factor out argument parsing and invocation of the `sudoku` package

## 2025-12-07

- All tests are passing
- I use `sudoku.New` to detect issues with the provided puzzle string
- I copied in my brute force sudoku solver as a fake implementation for `solve_sudoku`
  - this doesn't interoperate with the `sudoku` package at all

- [x] remove redundant `IsValidSudoku` check
  - [x] `sudoku.New` should return `ErrInvalid` for puzzles/2.invalid without calling `IsValidSudoku`

- [ ] implement a command that gets the next step towards solving the puzzle
  - e.g. if you provide a puzzle with a naked single in one position, the command prints something like `naked single in rXcY` and maybe the next state after filling in the naked single
  - internally we could have a unit testable interface that returns a struct defining the next action to take

---

# Scratch

- [ ] expect output to contain "naked single" when there's a naked single
- [ ] enable using the puzzles from `./tests/puzzles` as Go test fixtures to avoid duplication
  - create a package that embeds the puzzles directory
- [ ] instead of requiring puzzles to be provided as an 81 char string, make `solve_sudoku` handle ignoring whitespace and comments
  - [ ] accept puzzle on stdin
