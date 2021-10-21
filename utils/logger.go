package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// Logger Interface
type Logger interface {
	Info(msg string)
	Warn(msg string)
	Fatal(msg string)
	Debug(msg string)
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
}

type logger struct {
	Logger zerolog.Logger
}

// NewLogger creates a new logger instance.
func NewLogger(output *os.File, component string) (Logger, error) {
	log := zerolog.New(output).With().
		Str("component", component).
		Logger()

	switch viper.GetString("log-level") {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "warning":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msgf("Unknown log-level %s, using info.", viper.GetString("log-level"))
	}

	return logger{
		Logger: log,
	}, nil
}

func (l logger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l logger) Infof(msg string, args ...interface{}) {
	l.Logger.Info().Msgf(msg, args...)
}

func (l logger) Warn(msg string) {
	l.Logger.Warn().Msg(msg)
}

func (l logger) Warnf(msg string, args ...interface{}) {
	l.Logger.Warn().Msgf(msg, args...)
}

func (l logger) Fatal(msg string) {
	l.Logger.Fatal().Msg(msg)
}

func (l logger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

func (l logger) Fatalf(msg string, args ...interface{}) {
	l.Logger.Fatal().Msgf(msg, args...)
}

func (l logger) Debugf(msg string, args ...interface{}) {
	l.Logger.Debug().Msgf(msg, args...)
}
