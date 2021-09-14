package log

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/rapidashorg/temp-shared-libraries/env"
	"github.com/rs/zerolog"
)

// Config defines our custom logger config
type Config struct {
	AppName     string
	LogPath     string
	Level       string
	Environment env.Env

	ErrorWrapperStackTrace    bool
	ErrorWrapperExcludedCodes []int

	errorWrapperExcludedCodesMap map[int]bool
}

var (
	logger *zerolog.Logger
	config *Config
)

// openFile creates directories with -p and open LogPath file
func (c *Config) openFile() (*os.File, error) {
	if c.LogPath == "" {
		return nil, nil
	}

	err := os.MkdirAll(filepath.Dir(c.LogPath), 0755)
	if err != nil && err != os.ErrExist {
		return nil, err
	}

	return os.OpenFile(c.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}

// InitLogger initialize our custom logger
func InitLogger(tConfig *Config) error {
	var w io.Writer = os.Stderr
	if tConfig.Environment.IsDevelopment() {
		w = zerolog.ConsoleWriter{
			Out:     os.Stdout,
			NoColor: false,
		}
	}

	if file, err := tConfig.openFile(); err != nil {
		return err
	} else if file != nil {
		w = file
	}

	var level zerolog.Level
	switch tConfig.Level {
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "debug":
		level = zerolog.DebugLevel
	default:
		level = zerolog.ErrorLevel
	}

	zerolog.LevelFieldName = "lvl"
	zerolog.MessageFieldName = "msg"
	zerolog.TimestampFieldName = "ts"

	l := zerolog.New(w).
		With().Str("app", tConfig.AppName).Logger().
		Level(level)
	logger = &l

	config = tConfig
	config.errorWrapperExcludedCodesMap = make(map[int]bool)
	for _, v := range config.ErrorWrapperExcludedCodes {
		config.errorWrapperExcludedCodesMap[v] = true
	}

	return nil
}

func addGlobalLogEntries(ctx context.Context, e *zerolog.Event) *zerolog.Event {
	return e.
		Timestamp().
		Str("req_id", GetRequestID(ctx))
}
