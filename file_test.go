package testhelper

import (
	"os"
	"strings"
	"testing"
)

func TestTempdir(t *testing.T) {
	path, cleanup := Tempdir(t)
	if _, err := os.Stat(path); err != nil {
		t.Fail()
	}
	cleanup()
	if _, err := os.Stat(path); err == nil {
		t.Fail()
	}
}

func TestTempFile(t *testing.T) {
	t.Run("within subdir", func(t *testing.T) {
		td, tdCleanup := Tempdir(t)
		defer tdCleanup()
		path, cleanup := TempFile(t, td)
		if !strings.HasPrefix(path, td) {
			t.Fail()
		}
		if _, err := os.Stat(path); err != nil {
			t.Fail()
		}
		cleanup()
		if _, err := os.Stat(path); err == nil {
			t.Fail()
		}
	})

	t.Run("within subdir", func(t *testing.T) {
		path, cleanup := TempFile(t)
		if _, err := os.Stat(path); err != nil {
			t.Fail()
		}
		cleanup()
		if _, err := os.Stat(path); err == nil {
			t.Fail()
		}
	})
}

func TestTempFileExec(t *testing.T) {
	path, cleanup := TempFileExec(t)
	if _, err := os.Stat(path); err != nil {
		t.Fail()
	}
	cleanup()
	if _, err := os.Stat(path); err == nil {
		t.Fail()
	}
}
