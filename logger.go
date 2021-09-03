package phuslulogger

import (
	"os"
	"path/filepath"

	"github.com/phuslu/log"
)

const (
	DEFAULT_LOGFILE = "logs/log.log"
)

// for set log.Logger default
func SetDefaultlogger() {
	log.DefaultLogger = log.Logger{
		Level:  log.DebugLevel,
		Caller: 1,
		Writer: NewConsoleColor(),
	}
}

// // console log with text format
func GetLoggerConsole(logLevel log.Level, caller int) log.Logger {
	return log.Logger{
		Level:  logLevel,
		Caller: caller,
		Writer: NewConsoleColor(),
	}
}

// // console log with json format
func GetLoggerConsoleJson(logLevel log.Level, caller int) log.Logger {
	return log.Logger{
		Level:     logLevel,
		Caller:    caller,
		TimeField: "timestamp",
		Writer:    NewConsoleJson(),
	}
}

// // log file with json format
func GetLoggerFile(filelogName string, logLevel log.Level, caller int) log.Logger {
	er := os.MkdirAll(filepath.Dir(filelogName), 0755)
	if er != nil {
		log.Fatal().Err(er)
		os.Exit(1)
	}
	log.Debug().Msgf("logfile: %s", filelogName)
	DefaultFileWriter := NewFileWriter()
	DefaultFileWriter.Filename = filelogName
	logger := log.Logger{
		Level:  logLevel,
		Caller: caller,
		Writer: DefaultFileWriter,
	}
	return logger
}

// // log file with json format
func GetLoggerFileAndConsole(filelogName string, fileErrName string, logLevel log.Level, caller int) log.Logger {
	if er := os.MkdirAll(filepath.Dir(filelogName), 0755); er != nil {
		log.Fatal().Err(er)
		os.Exit(1)
	}
	if er := os.MkdirAll(filepath.Dir(fileErrName), 0755); er != nil {
		log.Fatal().Err(er)
		os.Exit(1)
	}
	log.Debug().Msgf("logfile: %s, errfile: %s", filelogName, fileErrName)
	errWriter := NewFileWriter()
	errWriter.Filename = fileErrName

	DefaultFileWriter := NewFileWriter()
	DefaultFileWriter.Filename = filelogName
	logger := log.Logger{
		Level:  logLevel,
		Caller: caller,
		Writer: &log.MultiWriter{
			ErrorWriter:   errWriter,
			WarnWriter:    DefaultFileWriter,
			InfoWriter:    DefaultFileWriter,
			ConsoleWriter: NewConsoleJson(),
			ConsoleLevel:  logLevel,
		},
	}
	return logger
}
