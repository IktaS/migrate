package migrate

import (
	nurl "net/url"
	"testing"
)

func TestSuintPanicsWithNegativeInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected suint to panic for -1")
		}
	}()
	suint64(-1)
}

func TestSuint(t *testing.T) {
	if u := suint64(0); u != 0 {
		t.Fatalf("expected 0, got %v", u)
	}
}

func TestFilterCustomQuery(t *testing.T) {
	n, err := nurl.Parse("foo://host?a=b&x-custom=foo&c=d&ok=y")
	if err != nil {
		t.Fatal(err)
	}
	nx := FilterCustomQuery(n).Query()
	if nx.Get("x-custom") != "" {
		t.Fatalf("didn't expect x-custom")
	}
	if nx.Get("ok") != "y" {
		t.Fatalf("expected ok=y, got %v", nx.Get("ok"))
	}
}
