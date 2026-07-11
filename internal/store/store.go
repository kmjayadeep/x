package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Store struct {
	path string
}

func New(name string) *Store {
	dir, err := os.UserCacheDir()
	if err != nil {
		dir = os.TempDir()
	}
	return &Store{path: filepath.Join(dir, "x", name)}
}

func (s *Store) Get(key string) string {
	values := s.load()
	return values[key]
}

func (s *Store) Set(key, value string) error {
	values := s.load()
	values[key] = value
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return err
	}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o600)
}

func (s *Store) load() map[string]string {
	values := make(map[string]string)
	data, err := os.ReadFile(s.path)
	if err == nil {
		_ = json.Unmarshal(data, &values)
	}
	return values
}
