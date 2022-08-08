package phuslulogger

import (
	"os"

	"github.com/phuslu/log"
)

var (
	listFileWriter []*log.FileWriter = nil
)

func NewFileWriter() *log.FileWriter {
	l := new(log.FileWriter)
	l.FileMode = 0644
	l.EnsureFolder = true
	l.MaxSize = 100 * 1024 * 1024
	l.LocalTime = true
	l.MaxBackups = 90
	l.TimeFormat = "20060102"
	l.ProcessID = true

	return l
}

func NewConsoleJson() *log.IOWriter {
	l := new(log.IOWriter)
	l.Writer = os.Stdout

	return l
}

func NewConsoleColor() *log.ConsoleWriter {
	l := new(log.ConsoleWriter)
	l.ColorOutput = true
	l.EndWithMessage = true
	return l
}
