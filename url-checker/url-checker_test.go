package urlchecker

import (
	"net"
	"regexp"
	"testing"
)

func Test_urlChecker_Validate(t *testing.T) {
	type fields struct {
		allowIP            bool
		reBlockedHosts     []*regexp.Regexp
		blockedCIDRs       []*net.IPNet
		reWhitelistedHosts []*regexp.Regexp
		whitelistedCIDRs   []*net.IPNet
	}
	type args struct {
		rawURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error parse url",
			args: args{
				rawURL: ":",
			},
			wantErr: true,
		},
		{
			name: "success host is whitelisted",
			fields: fields{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`0.0.0.0/0`)
						return r
					}(),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
			args: args{
				rawURL: "https://ecs7.tokopedia.com/internal-image.jpg",
			},
		},
		{
			name: "success ip is whitelisted",
			fields: fields{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`0.0.0.0/0`)
						return r
					}(),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
			args: args{
				rawURL: "https://10.1.16.22:8000/internal-image.jpg",
			},
		},
		{
			name: "error host is blocked",
			fields: fields{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`0.0.0.0/0`)
						return r
					}(),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
			args: args{
				rawURL: "https://internal.tokopedia.com/internal-image.jpg",
			},
			wantErr: true,
		},
		{
			name: "error ip is not allowed",
			fields: fields{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`0.0.0.0/0`)
						return r
					}(),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
			args: args{
				rawURL: "https://10.0.0.1/internal-image.jpg",
			},
			wantErr: true,
		},
		{
			name: "error ip is blocked",
			fields: fields{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				allowIP: true,
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`0.0.0.0/0`)
						return r
					}(),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
			args: args{
				rawURL: "https://127.0.0.1/internal-image.jpg",
			},
			wantErr: true,
		},
		{
			name: "success normal",
			fields: fields{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`0.0.0.0/0`)
						return r
					}(),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
			args: args{
				rawURL: "https://cf.shopee.com/image.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &urlChecker{
				allowIP:            tt.fields.allowIP,
				reBlockedHosts:     tt.fields.reBlockedHosts,
				blockedCIDRs:       tt.fields.blockedCIDRs,
				reWhitelistedHosts: tt.fields.reWhitelistedHosts,
				whitelistedCIDRs:   tt.fields.whitelistedCIDRs,
			}
			if err := h.Validate(tt.args.rawURL); (err != nil) != tt.wantErr {
				t.Errorf("urlChecker.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
