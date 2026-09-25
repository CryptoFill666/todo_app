package core_logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger

	file *os.File
}

func NewLogger() (*Logger, error) {
	file, err := os.OpenFile(
		"logs.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := slog.New(handler)

	return &Logger{
		Logger: logger,
		file:   file,
	}, nil
}
