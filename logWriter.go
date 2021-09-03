package phuslulogger

import (
	"os"

	"github.com/phuslu/log"
)

func NewFileWriter() *log.FileWriter {
	l := log.FileWriter{
		FileMode:     0644,
		MaxSize:      1000 * 1024 * 1024, // 1000MB
		EnsureFolder: true,
		LocalTime:    true,
		MaxBackups:   90, // 90 day
		TimeFormat:   "2006-01-02",
	}
	return &l
}

func NewConsoleJson() *log.IOWriter {
	l := log.IOWriter{os.Stdout}

	return &l
}

func NewConsoleColor() *log.ConsoleWriter {
	l := log.ConsoleWriter{
		ColorOutput:    true,
		EndWithMessage: true,
	}
	return &l
}
