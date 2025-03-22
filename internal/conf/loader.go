package conf

import (
	"errors"
	"fmt"
	"os"
	"sigs.k8s.io/yaml"
)

type Loader struct {
}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadConf(path string) (*RootConf, error) {
	if path == "" {
		return nil, errors.New("no conf file specified")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("conf file does not exist: %s", path)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read conf file: %w", err)
	}

	var rootConf *RootConf
	err = yaml.Unmarshal(file, &rootConf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse conf file: %w", err)
	}

	return rootConf, nil
}
