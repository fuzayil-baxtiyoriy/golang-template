package logger

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Log struct {
	zerolog.Logger
}

const (
	InfoLvl  = "info"
	DebugLvl = "debug"
	ErrorLvl = "error"
	TextFormat = "text"
	JsonFormat = "json"
	
)

var (
	ErrIncorrectLogLevel  = errors.New("incorrect log level")
	ErrIncorrectLogFormat = errors.New("incorrect log format")
	levelMapper           = map[string]zerolog.Level{
		DebugLvl: zerolog.DebugLevel,
		InfoLvl:  zerolog.InfoLevel,
		ErrorLvl: zerolog.ErrorLevel,
	}
)

func NewLogger(logLvl string, logFormat string) (*Log, error) {
	zerolog.TimeFieldFormat = time.RFC3339

	logLevel, ok := levelMapper[logLvl]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrIncorrectLogLevel, logLvl)
	}
	zerolog.SetGlobalLevel(logLevel)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	switch logFormat {
	case TextFormat:
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	case JsonFormat:
		// Default is JSON format
		// pass
	default:
		return nil, fmt.Errorf("%w: %s", ErrIncorrectLogFormat, logFormat)
	}

	return &Log{logger}, nil
}

func (l *Log) AddContext(key string, value any) *Log {
	return &Log{
		Logger: l.Logger.With().Interface(key, value).Logger(),
	}
}

func (l *Log) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

func (l *Log) Debugf(format string, v ...any) {
	l.Logger.Debug().Msgf(format, v...)
}

func (l *Log) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l *Log) Infof(format string, v ...any) {
	l.Logger.Info().Msgf(format, v...)
}

func (l *Log) Error(err error) {
	l.Logger.Err(err).Msg("")
}

func (l *Log) Errorf(format string, v ...any) {
	l.Logger.Err(fmt.Errorf(format, v...)).Msg("")
}

func (l *Log) Fatal(err error) {
	l.Logger.Fatal().Err(err).Msg("")
}

func (l *Log) Fatalf(format string, v ...any) {
	l.Logger.Fatal().Err(fmt.Errorf(format, v...)).Msg("")
}
