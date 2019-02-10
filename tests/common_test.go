package tests

import (
	"testing"

	"github.com/playground/httphash/common"
)

type Test struct {
	in  string
	out string
}

func TestHash(t *testing.T) {
	tests := []Test{
		Test{
			in:  "test",
			out: "098f6bcd4621d373cade4e832627b4f6",
		},
	}

	for _, test := range tests {
		out := common.MD5Hash(test.in)
		if out != test.out {
			t.Errorf("got '%s' want '%s'", out, test.out)
		}
	}
}

func TestURL(t *testing.T) {
	tests := []Test{
		Test{
			in:  "test",
			out: "http://test",
		},
		Test{
			in:  "http://google.com",
			out: "http://google.com",
		},
		Test{
			in:  "htp://google.com",
			out: "http://htp://google.com",
		},
		Test{
			in:  "www.google.com",
			out: "http://www.google.com",
		},
	}

	for _, test := range tests {
		out := common.ResolveURL(test.in)
		if out != test.out {
			t.Errorf("got '%s' want '%s'", out, test.out)
		}
	}
}
