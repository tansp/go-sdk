package ecloud

import (
	"net/http"

	"ecloud_gosdk.v1/auth"
	"ecloud_gosdk.v1/conf"
	"ecloud_gosdk.v1/lib/rpc"
)

type Config struct {
	AccessKey string
	SecretKey string
	Scheme    string
	MgHost    string
	UpHost    string
	Transport http.RoundTripper
}

type Client struct {
	rpc.Client
	mac *auth.Mac
	Config
}

func New(zone int, cfg *Config) (p *Client) {

	p = new(Client)
	if cfg != nil {
		p.Config = *cfg
	}

	p.mac = auth.NewMac(p.AccessKey, p.SecretKey)
	p.Client = rpc.Client{auth.NewClient(p.mac, p.Transport)}

	if zone < 0 || zone >= len(conf.Zones) {
		return
	}
	if p.UpHost == "" || p.MgHost == "" {
		p.MgHost = conf.Zones[zone].MgHosts[0]
		p.UpHost = conf.Zones[zone].UpHosts[0]
	}

	return
}

// 设置全局默认的 ACCESS_KEY, SECRET_KEY 变量。
func SetMac(accessKey, secretKey string) {

	conf.ACCESS_KEY, conf.SECRET_KEY = accessKey, secretKey
}

// 设置使用这个SDK的应用程序名。userApp 必须满足 [A-Za-z0-9]
func SetAppName(userApp string) error {

	return conf.SetAppName(userApp)
}
