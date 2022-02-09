package urlchecker

import (
	"net"
	"net/url"
)

func (h *urlChecker) Validate(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	hostStr := u.Hostname()
	host := []byte(hostStr)
	ip := net.ParseIP(hostStr)

	for _, v := range h.reWhitelistedHosts {
		if v.Match(host) {
			return nil
		}
	}

	if ip != nil {
		for _, v := range h.whitelistedCIDRs {
			if v.Contains(ip) {
				return nil
			}
		}
	}

	// skip checking blocked host and ip for force whitelist only
	if h.forceWhitelistOnly {
		return ErrHostBlocked
	}

	for _, v := range h.reBlockedHosts {
		if v.Match(host) {
			return ErrHostBlocked
		}
	}

	if ip != nil {
		if !h.allowIP {
			return ErrHostBlocked
		}

		for _, v := range h.blockedCIDRs {
			if v.Contains(ip) {
				return ErrHostBlocked
			}
		}
	}

	return nil
}
