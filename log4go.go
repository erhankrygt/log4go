package log4go

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type ILog interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

var (
	ErrFileCannotCreate = errors.New("log file cannot crate")
)

const (
	DEBUG = "DEBUG "
	INFO  = "INFO "
	WARN  = "WARN "
	ERROR = "ERROR "
	FATAL = "FATAL "
)

type Config struct {
	FileName string `env:"FILE_NAME"`
}

// compile-time proofs of log interface implementation
var _ ILog = (*Log)(nil)

type Log struct {
	file *os.File
}

func NewLog(c Config, env string) ILog {
	fn := fmt.Sprintf("%s_%s.log", c.FileName, env)
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(ErrFileCannotCreate)
	}
	return &Log{
		file: f,
	}
}

func (l Log) Debug(message string) {
	l.writer(DEBUG, message)
}

func (l Log) Info(message string) {
	l.writer(INFO, message)
}

func (l Log) Warn(message string) {
	l.writer(WARN, message)
}

func (l Log) Error(message string) {
	l.writer(ERROR, message)
}

func (l Log) Fatal(message string) {
	l.writer(FATAL, message)
}

func (l Log) writer(t string, m string) {
	logger := log.New(l.file, t, log.LstdFlags)
	logger.Println(m)
}
