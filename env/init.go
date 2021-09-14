package env

import (
	"net"
	"os"
	"runtime"
)

const (
	EnvKey = "TKPENV"

	DevelopmentEnv = "development"
	StagingEnv     = "staging"
	ProductionEnv  = "production"
)

type Env interface {
	GoVersion() string
	Hostname() string
	HostIP() string
	Env() string
	IsDevelopment() bool
	IsStaging() bool
	IsProduction() bool
}

type env struct {
	env       string
	hostname  string
	hostIP    string
	goVersion string
}

func New() Env {
	e := &env{
		hostname:  "-",
		hostIP:    "-",
		env:       DevelopmentEnv,
		goVersion: runtime.Version(),
	}

	if env := os.Getenv(EnvKey); env != "" && (env == DevelopmentEnv || env == StagingEnv || env == ProductionEnv) {
		e.env = env
	}

	if name, err := os.Hostname(); err == nil {
		e.hostname = name
	}

	if conn, err := net.Dial("udp", "8.8.8.8:80"); err == nil {
		defer conn.Close()

		if localAddr, ok := conn.LocalAddr().(*net.UDPAddr); ok {
			e.hostIP = localAddr.IP.String()
		}
	}

	return e
}
