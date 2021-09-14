package log

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	env_mock "github.com/rapidashorg/temp-shared-libraries/env/mock"
	"github.com/undefinedlabs/go-mpatch"
)

func TestInitLogger(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockEnv := env_mock.NewMockEnv(controller)
	mockEnv.EXPECT().IsDevelopment().Return(true).AnyTimes()

	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "TestInitLogger error mkdir",
			args: args{&Config{
				LogPath: "./",
			}},
			mock: func() {
				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(os.MkdirAll, func(path string, perm os.FileMode) error {
					defer p1.Unpatch()
					return errors.New("an error")
				})
				if err != nil {
					panic(err)
				}
			},
			wantErr: true,
		},
		{
			name: "TestInitLogger error open file",
			args: args{&Config{
				LogPath: "./",
			}},
			mock: func() {
				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(os.MkdirAll, func(path string, perm os.FileMode) error {
					defer p1.Unpatch()
					return nil
				})
				if err != nil {
					panic(err)
				}

				var p2 *mpatch.Patch
				p2, err = mpatch.PatchMethod(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
					defer p2.Unpatch()
					return nil, errors.New("an error")
				})
				if err != nil {
					panic(err)
				}
			},
			wantErr: true,
		},
		{
			name: "TestInitLogger success 1",
			args: args{&Config{
				LogPath: "",
			}},
		},
		{
			name: "TestInitLogger success 2",
			args: args{&Config{
				LogPath: "",
				Level:   "info",
			}},
		},
		{
			name: "TestInitLogger success 3",
			args: args{&Config{
				LogPath: "",
				Level:   "warn",
			}},
		},
		{
			name: "TestInitLogger success 4",
			args: args{&Config{
				LogPath: "",
				Level:   "error",
			}},
		},
		{
			name: "TestInitLogger success 5",
			args: args{&Config{
				LogPath: "",
				Level:   "debug",

				ErrorWrapperExcludedCodes: []int{100},
			}},
		},
		{
			name: "TestInitLogger success 6",
			args: args{&Config{
				LogPath: "./",
			}},
			mock: func() {
				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(os.MkdirAll, func(path string, perm os.FileMode) error {
					defer p1.Unpatch()
					return nil
				})
				if err != nil {
					panic(err)
				}

				var p2 *mpatch.Patch
				p2, err = mpatch.PatchMethod(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
					defer p2.Unpatch()
					return os.Stdout, nil
				})
				if err != nil {
					panic(err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			tt.args.config.Environment = mockEnv

			if err := InitLogger(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("InitLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
