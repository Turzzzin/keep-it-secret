package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Secret struct {
	Name      string            `json:"name"`
	Data      map[string]string `json:"data"`
	Encrypted string            `json:"encrypted"`
}

type Storage struct {
	mu      sync.Mutex
	secrets map[string]Secret
	file    string
}

func getSecretsFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".kis")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(dir, "secrets.json"), nil
}

func New() (*Storage, error) {
	file, err := getSecretsFile()
	if err != nil {
		return nil, err
	}

	s := &Storage{
		secrets: make(map[string]Secret),
		file:    file,
	}

	if _, err := os.Stat(file); err == nil {
		data, err := os.ReadFile(file)
		if err == nil && len(data) > 0 {
			_ = json.Unmarshal(data, &s.secrets)
		}
	}

	return s, nil
}

func (s *Storage) SaveSecret(name string, encrypted string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.secrets[name] = Secret{Name: name, Encrypted: encrypted}
	return s.persist()
}

func (s *Storage) GetSecret(name string) (Secret, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	secret, ok := s.secrets[name]
	if !ok {
		return Secret{}, errors.New("secret not found")
	}
	return secret, nil
}

func (s *Storage) ListSecrets() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	names := make([]string, 0, len(s.secrets))
	for name := range s.secrets {
		names = append(names, name)
	}
	return names
}

func (s *Storage) persist() error {
	data, err := json.MarshalIndent(s.secrets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.file, data, 0600)
}

func (s *Storage) DeleteSecret(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.secrets[name]; !ok {
		return errors.New("secret not found")
	}

	delete(s.secrets, name)
	return s.persist()
}

func (s *Storage) ClearAll() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.secrets = make(map[string]Secret)
	if err := os.Remove(s.file); err != nil && !os.IsNotExist(err) {
		return err
	}
	fmt.Println("⚠️  All secrets deleted.")
	return nil
}