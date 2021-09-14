test:
	@go test $(shell go list ./... | grep -Ev '/(mock)') -gcflags=-l -race -count=1 -cover

mockgen:
	@mockgen -source=httputil/httputil.go -destination=httputil/mock/httputil_mock.go -package=httputil_mock
	@mockgen -source=env/init.go -destination=env/mock/env_mock.go -package=env_mock
	@mockgen -source=monitoring/init.go -destination=monitoring/mock/monitoring_mock.go -package=monitoring_mock
	@mockgen -source=url-checker/init.go -destination=url-checker/mock/url-checker_mock.go -package=urlchecker_mock
