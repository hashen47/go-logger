package logger

import (
	"testing"
)

func TestTokenizeFormatString(t *testing.T) {
	type Testcase struct {
		formatString string
		tokens       *[]FormatterToken
		err          error
	}

	testcases := []Testcase{
		{
			"%datetime% %level%: %message%",
			&[]FormatterToken{
				{DATETIME, ""},
				{TEXT, " "},
				{LEVEL, ""},
				{TEXT, ": "},
				{MESSAGE, ""},
			},
			nil,
		},
		{
			"%datetime% %level%:",
			&[]FormatterToken{
				{DATETIME, ""},
				{TEXT, " "},
				{LEVEL, ""},
				{TEXT, ":"},
			},
			nil,
		},
		{
			"datetime level:",
			nil,
			&InvalidFormatStringError{"datetime level:", "at least should contains one other formatter token type (%message%, %level% or %datetime%)"},
		},
		{
			"	",
			nil,
			&InvalidFormatStringError{"	", "at least should contains one other formatter token type (%message%, %level% or %datetime%)"},
		},
		{
			"",
			nil,
			&InvalidFormatStringError{"", "format string cannot be empty"},
		},
		{
			" ",
			nil,
			&InvalidFormatStringError{" ", "at least should contains one other formatter token type (%message%, %level% or %datetime%)"},
		},
		{
			"%",
			nil,
			&InvalidFormatStringError{"%", "you only can use % sign with formatter token types"},
		},
		{
			"%%",
			nil,
			&InvalidFormatStringError{"%%", "formatter token type can only be %message%, %datetime% or %level%"},
		},
		{
			"%%%",
			nil,
			&InvalidFormatStringError{"%%%", "formatter token type can only be %message%, %datetime% or %level%"},
		},
		{
			"%%% % %          % %message%",
			nil,
			&InvalidFormatStringError{"%%% % %          % %message%", "formatter token type can only be %message%, %datetime% or %level%"},
		},
		{
			"%datetime% %level%: %something%",
			nil,
			&InvalidFormatStringError{"%datetime% %level%: %something%", "formatter token type can only be %message%, %datetime% or %level%"},
		},
		{
			"%datetime% %level%: %message",
			nil,
			&InvalidFormatStringError{"%datetime% %level%: %message", "you only can use % sign with formatter token types"},
		},
	}

	for _, tc := range testcases {
		tokens, err := tokenizeFormatString(tc.formatString)
		if tc.err != nil && err == nil ||
			tc.err == nil && err != nil {
			t.Fatalf("\nERROR\nFORMAT STRING: '%s'\nEXPECT: %v\nREAL: %v\n", tc.formatString, tc.err, err)
		}

		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("\nERROR\nFORMAT STRING: '%s'\nEXPECT: %v\nREAL: %v\n", tc.formatString, tc.err, err)
			}
		}

		if tokens == nil && tc.tokens != nil ||
			tokens != nil && tc.tokens == nil {
			t.Fatalf("\nERROR\nFORMAT STRING: '%s'\nEXPECT: %v\nREAL: %v\n", tc.formatString, tc.tokens, tokens)
		}

		if tokens != nil {
			if len(*tokens) != len(*tc.tokens) {
				t.Fatalf("\nERROR\nFORMAT STRING: '%s'\nEXPECT LEN: %d\nREAL LEN: %d\n", tc.formatString, len(*tc.tokens), len(*tokens))
			}

			for i := 0; i < len(*tokens); i++ {
				if (*tokens)[i].data != (*tc.tokens)[i].data {
					t.Fatalf("\nERROR\nFORMAT STRING: '%s'\nEXPECT DATA: '%s'\nREAL DATA: '%s'\n", tc.formatString, (*tc.tokens)[i].data, (*tokens)[i].data)
				}
			}
		}
	}
}
