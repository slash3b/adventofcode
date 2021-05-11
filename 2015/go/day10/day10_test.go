package main

import "testing"

func TestConvert(t *testing.T) {

	cached = make(map[string]string)

	cases := []struct {
		Incoming string
		Expected string
	}{
		{"1", "11"},
		{"2", "12"},
		{"3", "13"},
		{"11", "21"},
		{"111", "31"},
		{"22", "22"},
	}

	for _, c := range cases {
		if convert(c.Incoming) != c.Expected {
			t.Errorf("result mismatch incoming=%s expected=%s", c.Incoming, c.Expected)
		}
	}
}
