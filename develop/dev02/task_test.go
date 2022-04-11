package main

import (
	"errors"
	"testing"
)

type unpackTest struct {
	arg         string
	expected    string
	expectedErr error
}

func TestUnpack(t *testing.T) {
	var unpackTests = []unpackTest{
		{`a4bc2d5e`, "aaaabccddddde", nil},
		{`abcd`, "abcd", nil},
		{`45`, "", errors.New("incorrent first element")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "we44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}

	for _, test := range unpackTests {
		output, err := Unpack(test.arg)
		if output != test.expected && err != test.expectedErr {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
