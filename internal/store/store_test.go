package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	t.Parallel()

	s := &Store{path: filepath.Join(t.TempDir(), "nested", "state.json")}
	if got := s.Get("missing"); got != "" {
		t.Fatalf("Get(missing) = %q, want empty string", got)
	}
	if err := s.Set("duration", "30m"); err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	if got := s.Get("duration"); got != "30m" {
		t.Fatalf("Get(duration) = %q, want %q", got, "30m")
	}

	info, err := os.Stat(s.path)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Fatalf("state file permissions = %o, want 600", perm)
	}
}

func TestStorePreservesValues(t *testing.T) {
	t.Parallel()

	s := &Store{path: filepath.Join(t.TempDir(), "state.json")}
	if err := s.Set("first", "one"); err != nil {
		t.Fatal(err)
	}
	if err := s.Set("second", "two"); err != nil {
		t.Fatal(err)
	}
	if got := s.Get("first"); got != "one" {
		t.Fatalf("Get(first) = %q, want %q", got, "one")
	}
}
