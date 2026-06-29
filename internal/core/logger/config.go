package core_logger

import (
	"fmt"
	"os"
)

type Config struct {
	Level  string
	Folder string
}

func NewConfig() (Config, error) {
	level := os.Getenv("LOGGER_LEVEL")
	if level == "" {
		return Config{}, fmt.Errorf("LOGGER_LEVEL is required")
	}

	folder := os.Getenv("LOGGER_FOLDER")
	if folder == "" {
		return Config{}, fmt.Errorf("LOGGER_FOLDER is required")
	}

	return Config{
		Level:  level,
		Folder: folder,
	}, nil
}
