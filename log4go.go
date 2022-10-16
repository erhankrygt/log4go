package log4go

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type ILog interface {
	Debug(messages ...string)
	Info(messages ...string)
	Warn(messages ...string)
	Error(messages ...string)
	Fatal(messages ...string)
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

func (l Log) Debug(messages ...string) {
	l.writer(DEBUG, messages)
}

func (l Log) Info(messages ...string) {
	l.writer(INFO, messages)
}

func (l Log) Warn(messages ...string) {
	l.writer(WARN, messages)
}

func (l Log) Error(messages ...string) {
	l.writer(ERROR, messages)
}

func (l Log) Fatal(messages ...string) {
	l.writer(FATAL, messages)
}

func (l Log) writer(t string, m []string) {
	logger := log.New(l.file, t, log.LstdFlags)
	me := strings.Join(m, ",")
	logger.Println(me)
}
