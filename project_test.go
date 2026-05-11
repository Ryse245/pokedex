package main

import (
	"fmt"
	pokecache "pokedex/internal"
	"sync"
	"testing"
	"time"
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

func TestAddGet(t *testing.T) {
	mux := sync.Mutex{}
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval, &mux)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}

}
