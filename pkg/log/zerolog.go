package log

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Zerolog struct {
	logger zerolog.Logger
}

func NewZerolog(service string, version string) *Zerolog {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.LevelFieldName = "severity"
	zerolog.DurationFieldUnit = time.Microsecond
	zerolog.TimestampFieldName = "timestamp"

	zl := zerolog.
		New(os.Stderr).With().
		Str("service", service).
		Str("version", version).
		Timestamp().Logger()

	return &Zerolog{logger: zl}
}

func (zerolog *Zerolog) Trace(message string, fields Fields) {
	zerolog.logger.Trace().Fields(fields).Msg(message)
}

func (zerolog *Zerolog) Debug(message string, fields Fields) {
	zerolog.logger.Debug().Fields(fields).Msg(message)
}

func (zerolog *Zerolog) Info(message string, fields Fields) {
	zerolog.logger.Info().Fields(fields).Msg(message)
}

func (zerolog *Zerolog) Warn(message string, fields Fields) {
	zerolog.logger.Warn().Fields(fields).Msg(message)
}

func (zerolog *Zerolog) Error(message string, err error, fields Fields) {
	wrappedError := errors.Wrap(err, err.Error())
	zerolog.logger.Error().Stack().Err(wrappedError).Fields(fields).Msg(message)
}

func (zerolog *Zerolog) Fatal(message string, err error, fields Fields) {
	wrappedError := errors.Wrap(err, err.Error())
	zerolog.logger.Fatal().Stack().Err(wrappedError).Fields(fields).Msg(message)
}

func (zerolog *Zerolog) Panic(message string, err error, fields Fields) {
	wrappedError := errors.Wrap(err, err.Error())
	zerolog.logger.Panic().Stack().Err(wrappedError).Fields(fields).Msg(message)
}
