package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "d\t3s2v0", expected: "d\t\t\tss"},
		{input: "ab\"cd\"", expected: "ab\"cd\""},
		{input: "ab\\cd\\", expected: "ab\\cd\\"},
		{input: "\xff4\u00FF", expected: "����ÿ"},
		{input: "日本語", expected: "日本語"},
		{input: "\u65e52本\U00008a9e2", expected: "\u65e5\u65e5本\U00008a9e\U00008a9e"},
		{input: "rr-1", expected: "rr-"},
		{input: "%3@-", expected: "%%%@-"},
		{input: "в4ф0", expected: "вввв"},
		{input: "☀☁0☂3☃", expected: "☀☂☂☂☃"},
		{input: "☄★💪", expected: "☄★💪"},
		{input: "firs2t _1sec", expected: "firsst _sec"},
		{input: " ", expected: " "},
		{input: "f2 3s2", expected: "ff   ss"},
		{input: "ⅰr", expected: "ⅰr"},
		{input: "tⅦж", expected: "tⅦж"},
		{input: "⅓e", expected: "⅓e"},
		{input: "w⅓", expected: "w⅓"},
		{input: "🔥2⓲", expected: "🔥🔥⓲"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "\u65e522本\U00008a9e2", "4🤘"}
	for _, tc := range invalidStrings {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
