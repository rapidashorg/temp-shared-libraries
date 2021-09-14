package log

import (
	"context"
	"reflect"
	"testing"

	"github.com/rs/xid"
	"github.com/undefinedlabs/go-mpatch"
)

func TestInitLoggingContext(t *testing.T) {
	var p1 *mpatch.Patch
	p1, err := mpatch.PatchMethod(xid.New, func() xid.ID {
		return xid.ID{}
	})
	if err != nil {
		panic(err)
	}
	defer p1.Unpatch()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, "00000000000000000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitLoggingContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitLoggingContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setRequestID(t *testing.T) {
	var p1 *mpatch.Patch
	p1, err := mpatch.PatchMethod(xid.New, func() xid.ID {
		return xid.ID{}
	})
	if err != nil {
		panic(err)
	}
	defer p1.Unpatch()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "success; request id is already set",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyRequestID, "request-id"),
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, "request-id"),
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, "00000000000000000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setRequestID(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success empty",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyRequestID, "request-id"),
			},
			want: "request-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRequestID(tt.args.ctx); got != tt.want {
				t.Errorf("GetRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}
