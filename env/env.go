package env

func (e *env) GoVersion() string {
	return e.goVersion
}

func (e *env) Hostname() string {
	return e.hostname
}

func (e *env) HostIP() string {
	return e.hostIP
}

func (e *env) Env() string {
	return e.env
}

func (e *env) IsDevelopment() bool {
	return e.env == DevelopmentEnv
}

func (e *env) IsStaging() bool {
	return e.env == StagingEnv
}

func (e *env) IsProduction() bool {
	return e.env == ProductionEnv
}
