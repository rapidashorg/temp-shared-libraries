package env

import (
	"net"
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/undefinedlabs/go-mpatch"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		mock func()
		want Env
	}{
		{
			name: "success",
			mock: func() {
				os.Setenv(EnvKey, StagingEnv)

				var p1 *mpatch.Patch
				p1, err := mpatch.PatchMethod(runtime.Version, func() string {
					defer p1.Unpatch()
					return "go1.14"
				})
				if err != nil {
					panic(err)
				}

				var p2 *mpatch.Patch
				p2, err = mpatch.PatchMethod(os.Hostname, func() (string, error) {
					defer p2.Unpatch()
					return "server-hostname", nil
				})
				if err != nil {
					panic(err)
				}

				var p3 *mpatch.Patch
				p3, err = mpatch.PatchMethod(net.Dial, func(network string, address string) (net.Conn, error) {
					defer p3.Unpatch()
					return &net.UDPConn{}, nil
				})
				if err != nil {
					panic(err)
				}

				var p4 *mpatch.Patch
				p4, err = mpatch.PatchInstanceMethodByName(reflect.TypeOf(&net.UDPConn{}), "LocalAddr", func(c *net.UDPConn) net.Addr {
					defer p4.Unpatch()
					return &net.UDPAddr{
						IP: net.IPv4(1, 2, 3, 4),
					}
				})
				if err != nil {
					panic(err)
				}
			},
			want: &env{
				env:       "staging",
				hostname:  "server-hostname",
				hostIP:    "1.2.3.4",
				goVersion: "go1.14",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
