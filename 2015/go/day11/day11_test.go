package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestIsIncreasing(t *testing.T) {
	cases := []struct {
		incoming []byte
		expected bool
	}{
		{
			incoming: []byte("abcdddd"),
			expected: true,
		},
		{
			incoming: []byte("ddddabc"),
			expected: true,
		},
		{
			incoming: []byte("didooooooooooo"),
			expected: false,
		},
	}

	for _, c := range cases {
		if isIncreasing(c.incoming) != c.expected {
			t.Error("ooops", string(c.incoming))
		}
	}
}

func TestIsLegal(t *testing.T) {
	cases := []struct {
		incoming []byte
		expected bool
	}{
		{
			incoming: []byte("abcdddd"),
			expected: true,
		},
		{
			incoming: []byte("jjjjjjjjjjj_i"),
			expected: false,
		},
		{
			incoming: []byte("jjjjjjjjjjj_o"),
			expected: false,
		},
		{
			incoming: []byte("jjjjjjjjjjj_l"),
			expected: false,
		},
	}

	for _, c := range cases {
		if isLegalChars(c.incoming) != c.expected {
			t.Error("ooops", string(c.incoming))
		}
	}
}

func TestIsPairs(t *testing.T) {
	cases := []struct {
		incoming []byte
		expected bool
	}{
		{
			incoming: []byte("aabb"),
			expected: true,
		},
		{
			incoming: []byte("aaa"),
			expected: false,
		},
		{
			incoming: []byte("aakjlkjlkjljkljlbb"),
			expected: true,
		},
		{
			incoming: []byte("asdfasdfaabb"),
			expected: true,
		},
		{
			incoming: []byte("asdfgbbbjjasdgvas"),
			expected: true,
		},
		{
			incoming: []byte("qwertyuiop"),
			expected: false,
		},
	}

	for _, c := range cases {
		if isPairs(c.incoming) != c.expected {
			t.Error("ooops", string(c.incoming))
		}
	}
}

func TestRotation(t *testing.T) {
	testBytes := []byte("ay")
	flip(testBytes) // should be z
	flip(testBytes) // should be ba

	fmt.Println(testBytes)

	res := bytes.Compare(testBytes, []byte("ba"))
	if res != 0 {
		fmt.Println(testBytes)
		t.Error("nope")
	}
}
