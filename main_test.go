package main

import (
	"flag"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected struct {
			charset string
			length  int
			count   int
			ok      bool
		}
	}{
		{
			name: "no args",
			args: []string{},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "base64",
				length:  21,
				count:   1,
				ok:      true,
			},
		},
		{
			name: "canonic with count",
			args: []string{"10"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "base64",
				length:  21,
				count:   10,
				ok:      true,
			},
		},
		{
			name: "custom charset",
			args: []string{"456789", "10"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "456789",
				length:  10,
				count:   1,
				ok:      true,
			},
		},
		{
			name: "optional count",
			args: []string{"HEX", "10", "5"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "0123456789ABCDEF",
				length:  10,
				count:   5,
				ok:      true,
			},
		},
		{
			name: "wrong number of args",
			args: []string{"HEX"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "",
				length:  0,
				count:   0,
				ok:      false,
			},
		},
		{
			name: "invalid length",
			args: []string{"hex", "foo", "5"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "",
				length:  0,
				count:   0,
				ok:      false,
			},
		},
		{
			name: "out of bounds length",
			args: []string{"hex", "0", "5"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "",
				length:  0,
				count:   0,
				ok:      false,
			},
		},
		{
			name: "invalid count",
			args: []string{"hex", "10", "foo"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "",
				length:  0,
				count:   0,
				ok:      false,
			},
		},
		{
			name: "out of bounds count",
			args: []string{"hex", "10", "-1"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "",
				length:  0,
				count:   0,
				ok:      false,
			},
		},
		{
			name: "too many args",
			args: []string{"hex", "10", "10", "foo"},
			expected: struct {
				charset string
				length  int
				count   int
				ok      bool
			}{
				charset: "",
				length:  0,
				count:   0,
				ok:      false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(tt.name, flag.ExitOnError)
			err := flag.CommandLine.Parse(tt.args)
			if err != nil {
				t.Error(err)
			}
			charset, length, count, ok := parseFlags()
			if charset != tt.expected.charset || length != tt.expected.length || count != tt.expected.count || ok != tt.expected.ok {
				t.Errorf("parseFlags(%v) = (%v, %v, %v, %v), want (%v, %v, %v, %v)", tt.args, charset, length, count, ok, tt.expected.charset, tt.expected.length, tt.expected.count, tt.expected.ok)
			}
		})
	}
}
