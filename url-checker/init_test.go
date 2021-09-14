package urlchecker

import (
	"net"
	"reflect"
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name    string
		args    args
		want    URLChecker
		wantErr bool
	}{
		{
			name: "error compile blacklisted regex",
			args: args{
				cfg: &Config{
					BlacklistedHosts: []string{
						`(`,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error parse blacklisted cidr",
			args: args{
				cfg: &Config{
					BlacklistedHosts: []string{
						`.*\.?tokopedia\.com`,
					},
					BlacklistedCIDRs: []string{
						`256.256.256.256`,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error compile whitelisted regex",
			args: args{
				cfg: &Config{
					BlacklistedHosts: []string{
						`.*\.?tokopedia\.com`,
					},
					BlacklistedCIDRs: []string{
						`10.0.0.0/8`,
					},
					WhitelistedHosts: []string{
						`(`,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error parse whitelisted cidr",
			args: args{
				cfg: &Config{
					BlacklistedHosts: []string{
						`.*\.?tokopedia\.com`,
					},
					BlacklistedCIDRs: []string{
						`10.0.0.0/8`,
					},
					WhitelistedHosts: []string{
						`ecs7\.tokopedia\.com`,
					},
					WhitelistedCIDRs: []string{
						`256.256.256.256`,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				cfg: &Config{
					BlacklistedHosts: []string{
						`.*\.?tokopedia\.com`,
					},
					BlacklistedCIDRs: []string{
						`10.0.0.0/8`,
					},
					WhitelistedHosts: []string{
						`ecs7\.tokopedia\.com`,
					},
					WhitelistedCIDRs: []string{
						`10.1.0.0/16`,
					},
				},
			},
			want: &urlChecker{
				reBlockedHosts: []*regexp.Regexp{
					regexp.MustCompile(`.*\.?tokopedia\.com`),
				},
				blockedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.0.0.0/8`)
						return r
					}(),
				},
				reWhitelistedHosts: []*regexp.Regexp{
					regexp.MustCompile(`ecs7\.tokopedia\.com`),
				},
				whitelistedCIDRs: []*net.IPNet{
					func() *net.IPNet {
						_, r, _ := net.ParseCIDR(`10.1.0.0/16`)
						return r
					}(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
