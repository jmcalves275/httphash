package tests

import (
	"testing"

	"github.com/playground/httphash/client"
)

type TestClient struct {
	parallel int
	urls     []string
	out      *client.HTTPHash
}

type TestApp struct {
	parallel int
	urls     []string
	out      int
}

func TestResponseSize(t *testing.T) {
	tests := []TestApp{
		TestApp{
			parallel: 10,
			urls:     []string{"google.com"},
			out:      1,
		},
	}

	for _, test := range tests {
		httphash, err := client.New(test.parallel, test.urls)
		if err != nil {
			t.Errorf("got '%s' ", err)
		}

		out := httphash.Process()
		if out != test.out {
			t.Errorf("got '%+v' want '%+v'", out, test.out)
		}
	}

}

func TestNewHTTPHashClient(t *testing.T) {
	tests := []TestClient{
		TestClient{
			parallel: -1,
			urls:     []string{},
			out:      nil,
		},
	}

	for _, test := range tests {
		out, _ := client.New(test.parallel, test.urls)
		if out != test.out {
			t.Errorf("got '%+v' want '%+v'", out, test.out)
		}
	}
}
