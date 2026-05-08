package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for count, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of case %v actual does not match expected", count)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectWord := c.expected[i]
			if word != expectWord {
				t.Errorf("Word %v of case %v does not match expected", i, count)
				break
			}
		}
	}
}
