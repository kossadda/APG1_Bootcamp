package archiver

import (
	"os"
	"strings"
	"testing"
)

func TestRotateFiles(t *testing.T) {
	err := os.Mkdir(".test", 0755)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer os.RemoveAll(".test")

	args := []string{"-a", ".test", "archiver.go", "archiver_test.go"}
	arch, err := New(&args)
	if err != nil {
		t.Fatal(err)
	}

	arch.RotateFiles(args)

	info, err := os.ReadDir(".test")
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range info {
		if !strings.Contains(entry.Name(), "archiver") {
			t.Fatal("wrong filenames")
		}
	}
}

func TestRotateFail(t *testing.T) {
	err := os.Mkdir(".err", 0755)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer os.RemoveAll(".err")

	args := []string{"-a", ".err", "noexists"}
	arch, err := New(&args)
	if err != nil {
		t.Fatal(err)
	}

	arch.RotateFiles(args)

	info, err := os.ReadDir(".err")
	if err != nil {
		t.Fatal(err)
	}

	for range info {
		t.Fatal("no exceptions")
	}
}

func TestError(t *testing.T) {
	args := []string{"-a", ".test"}
	_, err := New(&args)
	if err == nil {
		t.Fatal(err)
	}
}
