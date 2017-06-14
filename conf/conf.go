package conf

import (
	"fmt"
	"runtime"

	"go-sdk/lib/rpc"
)

var version = "1.0"

var ACCESS_KEY string
var SECRET_KEY string

type ZoneConfig struct {
	UpHosts []string
	MgHosts []string
}

var Zones = []ZoneConfig{
	// z0 华北机房:
	{
		UpHosts: []string{

		//"http://192.168.200.12:8072",
		},
		MgHosts: []string{
		//	"http://192.168.100.98:80",
		},
	},
}

//userApp only can contains A-Z,a-z,0-9
func SetAppName(userApp string) error {

	rpc.UserAgent = fmt.Sprintf(
		"EcloudGo/%s (%s; %s; %s) %s", version, runtime.GOOS, runtime.GOARCH, userApp, runtime.Version())
	return nil
}

func init() {
	SetAppName("")
}
