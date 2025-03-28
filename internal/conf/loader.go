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

func (l *Loader) LoadConf(path string) (*ProjectsConf, error) {
	if path == "" {
		return nil, errors.New("no conf file specified")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("configuration file does not exist at path: %s", path)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read conf file: %w", err)
	}

	var rootConf *ProjectsConf
	err = yaml.Unmarshal(file, &rootConf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse conf file: %w", err)
	}

	return rootConf, nil
}
