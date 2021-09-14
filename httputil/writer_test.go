package httputil

import (
	"encoding/json"
	goerrors "errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/rapidashorg/temp-shared-libraries/errors"
	"github.com/undefinedlabs/go-mpatch"
)

func Test_httpUtil_WriteResponse(t *testing.T) {
	h := New()

	type args struct {
		w         http.ResponseWriter
		r         *http.Request
		startTime time.Time
		data      interface{}
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "error marshal",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
			mock: func() {
				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(json.Marshal, func(v interface{}) ([]byte, error) {
					defer p1.Unpatch()
					return nil, goerrors.New("an error")
				})
				if err != nil {
					panic(err)
				}
			},
		},
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			h.WriteResponse(tt.args.w, tt.args.r, tt.args.startTime, tt.args.data)
		})
	}
}

func Test_httpUtil_WriteErrorResponse(t *testing.T) {
	h := New()

	type args struct {
		w            http.ResponseWriter
		r            *http.Request
		err          *errors.ErrorWrapper
		startTime    time.Time
		errorContext string
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success",
			args: args{
				w:   httptest.NewRecorder(),
				r:   httptest.NewRequest("GET", "/", nil),
				err: &errors.ErrorWrapper{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			h.WriteErrorResponse(tt.args.w, tt.args.r, tt.args.err, tt.args.startTime, tt.args.errorContext)
		})
	}
}

func Test_httpUtil_WriteInternalServerError(t *testing.T) {
	h := New()

	type args struct {
		w           http.ResponseWriter
		processTime float64
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			h.WriteInternalServerError(tt.args.w, tt.args.processTime)
		})
	}
}

func Test_httpUtil_marshalJSONResponse(t *testing.T) {
	h := &httpUtil{}

	type args struct {
		processTime float64
		messages    []string
		data        interface{}
		statusCode  int
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []byte
		wantErr bool
	}{
		{
			name: "error marshal",
			mock: func() {
				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(json.Marshal, func(v interface{}) ([]byte, error) {
					defer p1.Unpatch()
					return nil, goerrors.New("an error")
				})
				if err != nil {
					panic(err)
				}
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				statusCode: 200,
			},
			mock: func() {
				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(json.Marshal, func(v interface{}) ([]byte, error) {
					defer p1.Unpatch()
					return []byte{}, nil
				})
				if err != nil {
					panic(err)
				}
			},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			got, err := h.marshalJSONResponse(tt.args.processTime, tt.args.messages, tt.args.data, tt.args.statusCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpUtil.marshalJSONResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpUtil.marshalJSONResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
