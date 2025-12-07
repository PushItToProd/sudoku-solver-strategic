package sudoku_test

import (
	"strings"
	"testing"

	"github.com/pushittoprod/sudoku-solver-strategic/pkg/sudoku"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		desc           string
		puzzle         string
		expectErr      bool
		expectedErrMsg string
	}{
		{
			desc:           "too short",
			puzzle:         "123",
			expectErr:      true,
			expectedErrMsg: "too short",
		},
		{
			desc:           "too long",
			puzzle:         "0389564177562149384913872566857913423496281751274356895621738948145697239738425619",
			expectErr:      true,
			expectedErrMsg: "too long",
		},
		{
			desc:      "valid and unsolved",
			puzzle:    "038956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expectErr: false,
		},
		{
			desc:           "invalid and unsolved",
			puzzle:         "028956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expectErr:      true,
			expectedErrMsg: "invalid puzzle",
		},
		{
			desc:           "invalid and complete",
			puzzle:         "228956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expectErr:      true,
			expectedErrMsg: "invalid puzzle",
		},
		{
			desc:           "fully solved",
			puzzle:         "238956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expectErr:      true,
			expectedErrMsg: "already solved",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := sudoku.New(tc.puzzle)

			if tc.expectErr {
				if err == nil {
					if tc.expectedErrMsg != "" {
						t.Errorf("expected error containing %q but got (nil)", tc.expectedErrMsg)
					} else {
						t.Error("expected error but got nil")
					}
				} else if gotErrMsg := err.Error(); tc.expectedErrMsg != "" && !strings.Contains(gotErrMsg, tc.expectedErrMsg) {
					t.Errorf("expected error containing %q but got %q", tc.expectedErrMsg, gotErrMsg)
				}
			} else if !tc.expectErr && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

type expectedSudokuState struct {
	isValid    bool
	isComplete bool
	isCorrect  bool
	isSolved   bool
}

func (ess expectedSudokuState) equals(ss sudoku.SudokuState) bool {
	return ess.isValid == ss.IsValid() &&
		ess.isComplete == ss.IsComplete() &&
		ess.isCorrect == ss.IsCorrect() &&
		ess.isSolved == ss.IsSolved()
}

func TestCheck(t *testing.T) {
	testCases := []struct {
		desc     string
		puzzle   string
		expected expectedSudokuState
	}{
		{
			desc:     "too short",
			puzzle:   "123",
			expected: expectedSudokuState{},
		},
		{
			desc:     "too long",
			puzzle:   "0389564177562149384913872566857913423496281751274356895621738948145697239738425619",
			expected: expectedSudokuState{},
		},
		{
			desc:     "invalid characters",
			puzzle:   "zz8956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expected: expectedSudokuState{false, true, false, false},
		},
		{
			desc:     "valid and incomplete",
			puzzle:   "038956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expected: expectedSudokuState{true, false, true, false},
		},
		{
			desc:     "valid, incorrect, and incomplete",
			puzzle:   "028956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expected: expectedSudokuState{true, false, false, false},
		},
		{
			desc:     "valid, incorrect, and complete",
			puzzle:   "228956417756214938491387256685791342349628175127435689562173894814569723973842561",
			expected: expectedSudokuState{true, true, false, false},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := sudoku.Check(tc.puzzle); !tc.expected.equals(got) {
				t.Errorf("expected %+v but got %+v", tc.expected, got)
			}
		})
	}
}
