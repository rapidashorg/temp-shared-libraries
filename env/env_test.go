package env

import (
	"testing"
)

func Test_env_GoVersion(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: "go1.14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.GoVersion(); got != tt.want {
				t.Errorf("env.GoVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_env_Hostname(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: "server-hostname",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.Hostname(); got != tt.want {
				t.Errorf("env.Hostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_env_HostIP(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: "1.2.3.4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.HostIP(); got != tt.want {
				t.Errorf("env.HostIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_env_Env(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: "staging",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.Env(); got != tt.want {
				t.Errorf("env.Env() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_env_IsDevelopment(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.IsDevelopment(); got != tt.want {
				t.Errorf("env.IsDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_env_IsStaging(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.IsStaging(); got != tt.want {
				t.Errorf("env.IsStaging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_env_IsProduction(t *testing.T) {
	type fields struct {
		env       string
		hostname  string
		hostIP    string
		goVersion string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &env{
				env:       tt.fields.env,
				hostname:  tt.fields.hostname,
				hostIP:    tt.fields.hostIP,
				goVersion: tt.fields.goVersion,
			}
			if got := e.IsProduction(); got != tt.want {
				t.Errorf("env.IsProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}
