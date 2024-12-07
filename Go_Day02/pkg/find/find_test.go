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
		"../../test/foo/bar":                           {},
		"../../test/foo/bar/broken_sl -> [broken]":     {},
		"../../test/foo/bar/test.txt":                  {},
		"../../test/foo/fou":                           {},
		"../../test/foo/fou/bar -> ../../test/foo/bar": {},
		"../../test/foo/fou/keep":                      {},
		"../../test/foo/fou/keep/.gitkeep":             {},
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
		absRoot + "/bar":                            {},
		absRoot + "/bar/broken_sl -> [broken]":      {},
		absRoot + "/bar/test.txt":                   {},
		absRoot + "/fou":                            {},
		absRoot + "/fou/bar -> " + absRoot + "/bar": {},
		absRoot + "/fou/keep":                       {},
		absRoot + "/fou/keep/.gitkeep":              {},
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
	root, err := filepath.Abs("../../test/foo")
	if err != nil {
		t.Fatal(err)
	}

	prm, err := param.New("test", []string{root})
	if err != nil {
		t.Fatal(err)
	}
	sys, err := Scan(prm)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]struct{}{
		filepath.Join(root, "bar"):                                    {},
		filepath.Join(root, "bar/broken_sl -> [broken]"):              {},
		filepath.Join(root, "bar/test.txt"):                           {},
		filepath.Join(root, "fou"):                                    {},
		filepath.Join(root, "fou/bar -> "+filepath.Join(root, "bar")): {},
		filepath.Join(root, "fou/keep"):                               {},
		filepath.Join(root, "fou/keep/.gitkeep"):                      {},
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
