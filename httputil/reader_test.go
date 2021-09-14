package httputil

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_httpUtil_GetURLParam(t *testing.T) {
	h := New()

	type args struct {
		r    *http.Request
		name string
	}
	tests := []struct {
		name string
		h    *httpUtil
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				r: httptest.NewRequest("GET", "/", nil),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := h.GetURLParam(tt.args.r, tt.args.name); got != tt.want {
				t.Errorf("httpUtil.GetURLParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpUtil_GetQueryParam(t *testing.T) {
	h := New()

	type args struct {
		r    *http.Request
		name string
	}
	tests := []struct {
		name string
		h    *httpUtil
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				r:    httptest.NewRequest("GET", "/?foo=bar", nil),
				name: "foo",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := h.GetQueryParam(tt.args.r, tt.args.name); got != tt.want {
				t.Errorf("httpUtil.GetQueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpUtil_ReadBody(t *testing.T) {
	h := New()

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *httpUtil
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				r: httptest.NewRequest("GET", "/", nil),
			},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := h.ReadBody(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpUtil.ReadBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpUtil.ReadBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
