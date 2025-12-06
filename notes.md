
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

2025-12-06

- I have a set of functional tests written in bash that check basic behavior of the app
- I have unit tests for a file `validator.go`
  - don't have a clear interface right now -- I implemented a couple redundant functions
    - `New` creates a `BoardState` or returns an error
    - `IsValidSudoku` just checks if a sudoku is valid
    - `Check` takes a puzzle as a string and returns a `SudokuState` object
  - idea: convert `SudokuState` to a set of error types or to a set of bit flags -- e.g. `ErrInvalidLength`, `ErrInvalidChars`, `ErrIncomplete`, `ErrConflictingDigits`
  - remember, the goal of this is to be able to provide useful error messages at the command line