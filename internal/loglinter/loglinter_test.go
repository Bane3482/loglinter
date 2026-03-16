package loglinter_test

import (
	"os"
	"path/filepath"
	"testing"

	"loglinter/internal/loglinter"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, loglinter.Analyzer, "loglinter")
}
