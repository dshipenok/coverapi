// +build coverage

package main

import (
	"flag"
	"io"
	"os"
	"testing"
)

// TestMain is an alternate entry point for tests
func TestMain(m *testing.M) {
	flag.Parse() // need to call manually
	main()
	flushReport()
}

func flushReport() {
	// Redirect Stdout/err so the testing code doesn't output:
	//		testing: warning: no tests to run
	//		PASS
	//		coverage: 79.3% of statements in ./...
	oldstdout := os.Stdout
	oldstderr := os.Stderr
	os.Stdout, _ = os.Open(os.DevNull)
	os.Stderr, _ = os.Open(os.DevNull)

	testing.MainStart(&testDeps{}, nil, nil, nil).Run()

	// restore
	os.Stdout = oldstdout
	os.Stderr = oldstderr
}

type testDeps struct{}

func (d testDeps) MatchString(pat, str string) (bool, error)   { return false, nil }
func (d testDeps) StartCPUProfile(io.Writer) error             { return nil }
func (d testDeps) StopCPUProfile()                             {}
func (d testDeps) WriteHeapProfile(io.Writer) error            { return nil }
func (d testDeps) WriteProfileTo(string, io.Writer, int) error { return nil }
func (d testDeps) ImportPath() string                          { return "" }
func (d testDeps) StartTestLog(io.Writer)                      {}
func (d testDeps) StopTestLog() error                          { return nil }
