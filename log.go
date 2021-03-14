package log

import (
	"fmt"
	"io/ioutil"
	slog "log"
	"os"
)

type levelLog int

const (
	CommonLog levelLog = iota
	DebugLog
	TraceLog
)

var (
	Trace   *slog.Logger
	Debug   *slog.Logger
	Info    *slog.Logger
	Warning *slog.Logger
	Error   *slog.Logger
	Fatal   *slog.Logger
)

func Init(pathLog string, lvlLog levelLog, isTest bool) error{
	var err error
	var file    *os.File

	writer := ioutil.Discard

	if !isTest {
		file, err = os.OpenFile(pathLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("can't open log file: %v", err)
		}
		writer = file
	}

	Info = slog.New(writer,
		"INFO:    ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Warning = slog.New(writer,
		"WARNING: ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Error = slog.New(writer,
		"ERROR:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Fatal = slog.New(writer,
		"FATAL:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Debug = slog.New(ioutil.Discard,
		"DEBUG:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Trace = slog.New(ioutil.Discard,
		"TRACE:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	switch lvlLog {
	case DebugLog:
		Debug = slog.New(writer,
			"DEBUG:   ",
			slog.Ldate|slog.Ltime|slog.Lshortfile)
	case TraceLog:
		Debug = slog.New(writer,
			"DEBUG:   ",
			slog.Ldate|slog.Ltime|slog.Lshortfile)
		Trace = slog.New(writer,
			"TRACE:   ",
			slog.Ldate|slog.Ltime|slog.Lshortfile)
	default:
	}
	return nil
}
