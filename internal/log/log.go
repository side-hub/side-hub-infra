package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

const (
	Name           = "app.log"
	Level          = "trace"
	LogTimeFormat  = "2006-01-02 15:04:05.000"
	SkipFrameCount = 2
	ExpireTime     = 24 * time.Hour
	RotationTime   = time.Hour
)

func Init(path string) error {
	if path != "" {
		os.MkdirAll(path, os.ModePerm|os.ModeDir)
	}

	level, err := zerolog.ParseLevel(Level)
	if err != nil {
		return errors.Wrap(err, "invalid level")
	}

	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = LogTimeFormat

	if path != "" {
		info, err := os.Stat(path)
		if err == nil && info.IsDir() {
			return fileLog(path)
		}
	}

	return consoleLog()
}

func consoleLog() error {
	zlog.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: LogTimeFormat}).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + SkipFrameCount).
		Logger()

	Debug("log mode: stdout")
	return nil
}

func fileLog(parent string) error {
	path := filepath.Join(parent, Name)
	rawPath := filepath.Join(parent, "log", "%Y-%m-%d_%H:%M.log")

	options := []rotatelogs.Option{
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(RotationTime),
		rotatelogs.WithMaxAge(ExpireTime),
	}

	fileLogger, err := rotatelogs.New(rawPath, options...)
	if err != nil {
		return errors.Wrap(err, "failed to create rotatelogs")
	}

	// file, console multi writer
	zlog.Logger = zerolog.New(zerolog.MultiLevelWriter(fileLogger, zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: LogTimeFormat})).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + SkipFrameCount).
		Logger()

	Debugf("log mode: stdout, %s", path)

	return nil
}

func Debug(args ...any) {
	zlog.Debug().Msg(fmt.Sprint(args...))
}

func Info(args ...any) {
	zlog.Info().Msg(fmt.Sprint(args...))
}

func Warn(args ...any) {
	zlog.Warn().Msg(fmt.Sprint(args...))
}

func Error(args ...any) {
	zlog.Error().Msg(fmt.Sprint(args...))
}

func Fatal(args ...any) {
	zlog.Fatal().Msg(fmt.Sprint(args...))
}

func Panic(args ...any) {
	zlog.Panic().Msg(fmt.Sprint(args...))
}

func Trace(args ...any) {
	zlog.Trace().Msg(fmt.Sprint(args...))
}

func Debugf(format string, args ...any) {
	zlog.Debug().Msgf(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...any) {
	zlog.Info().Msgf(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	zlog.Warn().Msgf(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	zlog.Error().Msgf(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...any) {
	zlog.Fatal().Msgf(fmt.Sprintf(format, args...))
}

func Panicf(format string, args ...any) {
	zlog.Panic().Msgf(fmt.Sprintf(format, args...))
}

func Tracef(format string, args ...any) {
	zlog.Trace().Msgf(fmt.Sprintf(format, args...))
}
