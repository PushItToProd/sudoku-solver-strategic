
- [ ] expect output to contain "naked single" when there's a naked single

- [x] start unit tests
  - [ ] define internal interface


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


- `solve_sudoku` entrypoint
  - [ ] validation and error reporting - `solve_sudoku` should rely on `sudoku.New()` to report error states
  - [ ] factor out argument parsing and invocation of the `sudoku` package
- [ ] implement solving: use my bruteforce solver implementation to get output working for now
