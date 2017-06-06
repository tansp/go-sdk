package conf

import (
	"fmt"
	"runtime"

	"ecloud_gosdk.v1/lib/rpc"
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
		//"http://223.99.255.186:6051",
		//"http://upload.qiniu.com",
		//"http://upload.qiniu.com",
		},
		MgHosts: []string{
			"http://192.168.100.98:80",
			//"http://223.99.255.201:6057",
			//"http://223.99.255.212:6057",
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
