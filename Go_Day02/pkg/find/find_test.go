package find

import (
	"github.com/kossadda/APG1_Bootcamp/pkg/param"
	"path/filepath"
	"testing"
)

func TestScan(t *testing.T) {
	root := "../../test/foo"
	prm, err := param.New("test", []string{root})
	if err != nil {
		t.Fatal(err)
	}
	sys, err := Scan(prm)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]struct{}{
		"../../test/foo/bar":                                {},
		"../../test/foo/bar/baz":                            {},
		"../../test/foo/bar/baz/deep":                       {},
		"../../test/foo/bar/baz/deep/directory":             {},
		"../../test/foo/bar/broken_sl -> [broken]":          {},
		"../../test/foo/bar/buzz -> ../../test/foo/bar/baz": {},
		"../../test/foo/bar/test.txt":                       {},
	}

	for _, path := range sys {
		if _, ok := expected[path]; !ok {
			t.Errorf("unexpected path: %s", path)
		}
		delete(expected, path)
	}

	if len(expected) > 0 {
		for path := range expected {
			t.Errorf("missing expected path: %s", path)
		}
	}
}

func TestScanWithAbsolutePaths(t *testing.T) {
	absRoot, err := filepath.Abs("../../test/foo")
	if err != nil {
		t.Fatal(err)
	}

	prm, err := param.New("test", []string{absRoot})
	if err != nil {
		t.Fatal(err)
	}
	sys, err := Scan(prm)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]struct{}{
		absRoot + "/bar":                                 {},
		absRoot + "/bar/baz":                             {},
		absRoot + "/bar/baz/deep":                        {},
		absRoot + "/bar/baz/deep/directory":              {},
		absRoot + "/bar/broken_sl -> [broken]":           {},
		absRoot + "/bar/buzz -> " + absRoot + "/bar/baz": {},
		absRoot + "/bar/test.txt":                        {},
	}

	for _, path := range sys {
		if _, ok := expected[path]; !ok {
			t.Errorf("unexpected path: %s", path)
		}
		delete(expected, path)
	}

	if len(expected) > 0 {
		for path := range expected {
			t.Errorf("missing expected path: %s", path)
		}
	}
}

func TestScanEmptyDirectory(t *testing.T) {
	root := "../../test/empty"
	prm, err := param.New("test", []string{root})
	if err != nil {
		t.Fatal(err)
	}
	sys, err := Scan(prm)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]struct{}{}

	for _, path := range sys {
		if _, ok := expected[path]; !ok {
			t.Errorf("unexpected path: %s", path)
		}
		delete(expected, path)
	}

	if len(expected) > 0 {
		for path := range expected {
			t.Errorf("missing expected path: %s", path)
		}
	}
}

func TestScanFlexiblePaths(t *testing.T) {
	root := "../../test/foo"
	prm, err := param.New("test", []string{root})
	if err != nil {
		t.Fatal(err)
	}
	sys, err := Scan(prm)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]struct{}{
		filepath.Clean("../../test/foo/bar"):                                {},
		filepath.Clean("../../test/foo/bar/baz"):                            {},
		filepath.Clean("../../test/foo/bar/baz/deep"):                       {},
		filepath.Clean("../../test/foo/bar/baz/deep/directory"):             {},
		filepath.Clean("../../test/foo/bar/broken_sl -> [broken]"):          {},
		filepath.Clean("../../test/foo/bar/buzz -> ../../test/foo/bar/baz"): {},
		filepath.Clean("../../test/foo/bar/test.txt"):                       {},
	}

	for _, path := range sys {
		path = filepath.Clean(path)
		if _, ok := expected[path]; !ok {
			t.Errorf("unexpected path: %s", path)
		}
		delete(expected, path)
	}

	if len(expected) > 0 {
		for path := range expected {
			t.Errorf("missing expected path: %s", path)
		}
	}
}
