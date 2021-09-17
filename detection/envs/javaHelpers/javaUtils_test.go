package javaHelpers

import "testing"

type testPair struct {
	input    string
	expected string
}

func TestConvertJavaVersion(t *testing.T) {
	testCases := []testPair{
		{"1.8", "8"},
		{"1.9.3", "9.3"},
		{"10.0", "10.0"},
		{"abc", "abc"},
	}

	for i, testCase := range testCases {
		res, _ := ConvertJavaVersion(testCase.input)
		if res != testCase.expected {
			t.Errorf("(%d) got %s, expected %s", i, res, testCase.expected)
		}
	}
}
