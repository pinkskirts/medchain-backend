package core

import (
	"io"
	"log"
	"os"
)

const (
	FLAGS = log.LstdFlags | log.Lshortfile
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	fatal   *log.Logger
	panic   *log.Logger
}

func buildLoggerStruct(l *log.Logger) *Logger {
	return &Logger{
		debug:   log.New(l.Writer(), "DEBUG:", l.Flags()),
		info:    log.New(l.Writer(), "INFO:", l.Flags()),
		warning: log.New(l.Writer(), "WARNING:", l.Flags()),
		error:   log.New(l.Writer(), "ERROR:", l.Flags()),
		fatal:   log.New(l.Writer(), "FATAL:", l.Flags()),
		panic:   log.New(l.Writer(), "PANIC:", l.Flags()),
	}
}

// BuildStdoutLogger return a logger struct with writer direct to stdout.
func BuildStdoutLogger(s string) *Logger {
	logger := log.New(io.Writer(os.Stdout), s, FLAGS)

	return buildLoggerStruct(logger)
}

// BuildLogger return a logger struct with writer direct to w func param.
func BuildLogger(s string, w io.Writer) *Logger {
	logger := log.New(w, s, FLAGS)

	return buildLoggerStruct(logger)
}

// Debug calls logger Print (l.Output), with "DEBUG" flag in prefix.
func (l *Logger) Debug(args ...interface{}) {
	l.debug.Print(args)
}

// Info calls logger Print (l.Output), with "INFO" flag in prefix.
func (l *Logger) Info(args ...interface{}) {
	l.info.Print(args)
}

// Warning calls logger Print (l.Output), with "WARNING" flag in prefix.
func (l *Logger) Warning(args ...interface{}) {
	l.warning.Print(args)
}

// Error calls logger Print (l.Output), with "ERROR" flag in prefix.
func (l *Logger) Error(args ...interface{}) {
	l.error.Print(args)
}

// Fatal calls logger l.Print follow by os.Exit(1), with "FATAL" flag in prefix.
func (l *Logger) Fatal(args ...interface{}) {
	l.fatal.Fatal(args)
}

// Panic calls logger l.Print follow by panic(), with "PANIC" flag in prefix.
func (l *Logger) Panic(args ...interface{}) {
	l.panic.Panic(args)
}

// Debugf calls logger Printf (l.Output), with "DEBUG" flag in prefix.
// Arguments are handled in the manner of [fmt.Printf]
func (l *Logger) Debugf(f string, args ...interface{}) {
	l.debug.Printf(f, args)
}

// Infof calls logger Printf (l.Output), with "INFO" flag in prefix.
// Arguments are handled in the manner of [fmt.Printf]
func (l *Logger) Infof(f string, args ...interface{}) {
	l.info.Printf(f, args)
}

// Warningf calls logger Printf (l.Output), with "WARNING" flag in prefix.
// Arguments are handled in the manner of [fmt.Printf]
func (l *Logger) Warningf(f string, args ...interface{}) {
	l.warning.Printf(f, args)
}

// Errorf calls logger Printf (l.Output), with "ERROR" flag in prefix.
// Arguments are handled in the manner of [fmt.Printf]
func (l *Logger) Errorf(f string, args ...interface{}) {
	l.error.Printf(f, args)
}

// Fatal calls logger l.Printf follow by os.Exit(1), with "FATAL" flag in prefix.
func (l *Logger) Fatalf(f string, args ...interface{}) {
	l.fatal.Fatalf(f, args)
}

// Panic calls logger l.Printf follow by panic(), with "PANIC" flag in prefix.
func (l *Logger) Panicf(f string, args ...interface{}) {
	l.panic.Panicf(f, args)
}
