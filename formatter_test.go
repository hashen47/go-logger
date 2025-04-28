package logger

import (
	"testing"
)

func Test_NewFormatter(t *testing.T) {
	type Testcase struct {
		formatString string
		err          error
	}

	testcases := []Testcase{
		{
			"%message%",
			nil,
		},
		{
			"%level%",
			nil,
		},
		{
			"%datetime%",
			nil,
		},
		{
			"%level%%datetime%",
			nil,
		},
		{
			"%level%%datetime%%message%",
			nil,
		},
		{
			"%message% %message% %level% %datetime% %datetime% something here also",
			nil,
		},
		{
			"%level%datetime%",
			&InvalidFormatStringError{"%level%datetime%", "you only can use % sign with formatter token types"},
		},
		{
			"%level%%datetime",
			&InvalidFormatStringError{"%level%%datetime", "you only can use % sign with formatter token types"},
		},
		{
			"leveldatetime",
			&InvalidFormatStringError{"leveldatetime", "at least should contains one other formatter token type (%message%, %level% or %datetime%)"},
		},
		{
			"",
			&InvalidFormatStringError{"", "format string cannot be empty"},
		},
		{
			"%level%%something%",
			&InvalidFormatStringError{"%level%%something%", "formatter token type can only be %message%, %datetime% or %level%"},
		},
	}

	for _, tc := range testcases {
		_, err := _NewFormatter(tc.formatString)

		if err != nil && tc.err == nil ||
			err == nil && tc.err != nil {
			t.Fatalf("\nERROR\nEXPECT: %v\nREAL: %v\n", tc.err, err)
		}

		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("\nERROR\nEXPECT: %v\nREAL: %v\n", tc.err, err)
			}
		}
	}
}
