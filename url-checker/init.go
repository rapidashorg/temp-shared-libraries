package urlchecker

import (
	"net"
	"regexp"

	"github.com/pkg/errors"
)

var ErrHostBlocked = errors.New("url host is blocked")

type URLChecker interface {
	Validate(rawURL string) error
}

type urlChecker struct {
	allowIP bool

	reBlockedHosts []*regexp.Regexp
	blockedCIDRs   []*net.IPNet

	reWhitelistedHosts []*regexp.Regexp
	whitelistedCIDRs   []*net.IPNet
}

type Config struct {
	AllowIP bool

	BlacklistedHosts []string
	BlacklistedCIDRs []string

	WhitelistedHosts []string
	WhitelistedCIDRs []string
}

func New(cfg *Config) (URLChecker, error) {
	reBlockedHosts := make([]*regexp.Regexp, 0)
	for _, v := range cfg.BlacklistedHosts {
		r, err := regexp.Compile(v)
		if err != nil {
			return nil, errors.Wrapf(err, "url-checker.New.compile-blacklist-host-regex: %s", v)
		}

		reBlockedHosts = append(reBlockedHosts, r)
	}

	blockedCIDRs := make([]*net.IPNet, 0)
	for _, v := range cfg.BlacklistedCIDRs {
		_, in, err := net.ParseCIDR(v)
		if err != nil {
			return nil, errors.Wrapf(err, "url-checker.New.parse-blacklist-cidr: %s", v)
		}

		blockedCIDRs = append(blockedCIDRs, in)
	}

	reWhitelistedHosts := make([]*regexp.Regexp, 0)
	for _, v := range cfg.WhitelistedHosts {
		r, err := regexp.Compile(v)
		if err != nil {
			return nil, errors.Wrapf(err, "url-checker.New.compile-whitelist-host-regex: %s", v)
		}

		reWhitelistedHosts = append(reWhitelistedHosts, r)
	}

	whitelistedCIDRs := make([]*net.IPNet, 0)
	for _, v := range cfg.WhitelistedCIDRs {
		_, in, err := net.ParseCIDR(v)
		if err != nil {
			return nil, errors.Wrapf(err, "url-checker.New.parse-whitelist-cidr: %s", v)
		}

		whitelistedCIDRs = append(whitelistedCIDRs, in)
	}

	return &urlChecker{
		allowIP: cfg.AllowIP,

		reBlockedHosts: reBlockedHosts,
		blockedCIDRs:   blockedCIDRs,

		reWhitelistedHosts: reWhitelistedHosts,
		whitelistedCIDRs:   whitelistedCIDRs,
	}, nil
}
