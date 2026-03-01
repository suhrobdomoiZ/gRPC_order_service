package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type ConfigKey string

func (key ConfigKey) MustGet() string {
	val := os.Getenv(string(key))

	if val == "" {
		panic(fmt.Sprintf("config.MustGet: required value %s is not set", string(key)))
	}
	return val

}

func (key ConfigKey) Get(defaultValue string) string {
	if val := os.Getenv(string(key)); val != "" {
		return val
	}
	return defaultValue
}

func LoadDotEnv(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return fmt.Errorf("config.LoadDotEnv: %w", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")

		if !found {
			continue
		}
		key = strings.TrimSpace(key)

		if os.Getenv(key) == "" {
			if err := os.Setenv(key, value); err != nil {
				return fmt.Errorf("config.LoadDotEnv: %w", err)
			}
		}
	}
	return nil
}
