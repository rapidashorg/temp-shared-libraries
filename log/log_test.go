package log

import (
	"context"
	"testing"

	"github.com/rapidashorg/temp-shared-libraries/errors"
)

func TestInfo(t *testing.T) {
	loggerInst := logger

	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success no logger",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = nil
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = loggerInst
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			Info(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestWarn(t *testing.T) {
	loggerInst := logger

	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success no logger",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = nil
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = loggerInst
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			Warn(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestError(t *testing.T) {
	loggerInst := logger

	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success no logger",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = nil
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = loggerInst
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			Error(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}

func TestErrorWrapper(t *testing.T) {
	loggerInst := logger

	type args struct {
		ctx          context.Context
		err          *errors.ErrorWrapper
		errorContext string
	}
	tests := []struct {
		name string
		args args
		mock func()
	}{
		{
			name: "success no logger",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = nil
			},
		},
		{
			name: "success 1",
			args: args{
				ctx: context.Background(),
				err: &errors.ErrorWrapper{
					Code: 100,
				},
			},
			mock: func() {
				logger = loggerInst
				config = &Config{
					errorWrapperExcludedCodesMap: map[int]bool{
						100: true,
					},
				}
			},
		},
		{
			name: "success 2",
			args: args{
				ctx: context.Background(),
				err: &errors.ErrorWrapper{
					AdditionalData: []*errors.ErrorWrapperAdditionalData{
						{
							Line: "sample-line",
							Data: map[string]interface{}{
								"key": "value",
							},
						},
					},
				},
			},
			mock: func() {
				config = &Config{
					ErrorWrapperStackTrace: true,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			ErrorWrapper(tt.args.ctx, tt.args.err, tt.args.errorContext)
		})
	}
}

func TestDebug(t *testing.T) {
	loggerInst := logger

	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success no logger",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = nil
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				logger = loggerInst
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			Debug(tt.args.ctx, tt.args.format, tt.args.args...)
		})
	}
}
