package record

import (
	"testing"
)

var badPaths = []string{
	"foo/bar/baz",
	"//foo/bar/baz",
	"/ns",
	"ns",
	"ns/",
	"",
	"//",
	"/",
	"////",
}

func TestSplitPath(t *testing.T) {
	ns, key, err := SplitKey("/foo/bar/baz")
	if err != nil {
		t.Fatal(err)
	}
	if ns != "foo" {
		t.Errorf("wrong namespace: %s", ns)
	}
	if key != "bar/baz" {
		t.Errorf("wrong key: %s", key)
	}

	ns, key, err = SplitKey("/foo/bar")
	if err != nil {
		t.Fatal(err)
	}
	if ns != "foo" {
		t.Errorf("wrong namespace: %s", ns)
	}
	if key != "bar" {
		t.Errorf("wrong key: %s", key)
	}

	for _, badP := range badPaths {
		_, _, err := SplitKey(badP)
		if err == nil {
			t.Errorf("expected error for bad path: %s", badP)
		}
	}
}

func TestBestRecord(t *testing.T) {
	sel := NamespacedValidator{
		"pk": PublicKeyValidator{},
	}

	i, err := sel.Select("/pk/thing", [][]byte{[]byte("first"), []byte("second")})
	if err != nil {
		t.Fatal(err)
	}
	if i != 0 {
		t.Error("expected to select first record")
	}

	_, err = sel.Select("/pk/thing", nil)
	if err == nil {
		t.Fatal("expected error for no records")
	}

	_, err = sel.Select("/other/thing", [][]byte{[]byte("first"), []byte("second")})
	if err == nil {
		t.Fatal("expected error for unregistered ns")
	}

	_, err = sel.Select("bad", [][]byte{[]byte("first"), []byte("second")})
	if err == nil {
		t.Fatal("expected error for bad key")
	}
}
