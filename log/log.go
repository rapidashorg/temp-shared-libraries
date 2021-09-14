package log

import (
	"context"

	"github.com/rapidashorg/temp-shared-libraries/errors"
)

// Info logs a message with info level
func Info(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		return
	}

	addGlobalLogEntries(ctx, logger.Info()).
		Msgf(format, args...)
}

// Warn logs a message with warn level
func Warn(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		return
	}

	addGlobalLogEntries(ctx, logger.Warn()).
		Msgf(format, args...)
}

// Error logs a message with error level
func Error(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		return
	}

	addGlobalLogEntries(ctx, logger.Error()).
		Msgf(format, args...)
}

// ErrorWrapper logs an ErrorWrapper with error level
func ErrorWrapper(ctx context.Context, err *errors.ErrorWrapper, errorContext string) {
	if logger == nil {
		return
	}

	if err == nil || config.errorWrapperExcludedCodesMap[err.Code] {
		return
	}

	type additionalData struct {
		Line string                 `json:"l"`
		Data map[string]interface{} `json:"d"`
	}

	type errJSON struct {
		Context        string           `json:"ctx,omitempty"`
		StackTrace     []string         `json:"stk,omitempty"`
		AdditionalData []additionalData `json:"dta,omitempty"`
	}

	e := errJSON{
		Context: errorContext,
	}

	if config != nil && config.ErrorWrapperStackTrace {
		e.StackTrace = err.StackTrace
	}

	if err.AdditionalData != nil {
		e.AdditionalData = make([]additionalData, 0)
		for _, v := range err.AdditionalData {
			e.AdditionalData = append(e.AdditionalData, additionalData{
				Line: v.Line,
				Data: v.Data,
			})
		}
	}

	addGlobalLogEntries(ctx, logger.Error()).
		Interface("err", e).
		Msg(err.ActualError())
}

// Debug logs a message with debug level
func Debug(ctx context.Context, format string, args ...interface{}) {
	if logger == nil {
		return
	}

	addGlobalLogEntries(ctx, logger.Debug()).
		Msgf(format, args...)
}
