package ecloudcli

import (
	"net/http"

	"ecloud_gosdk.v1/conf"
	"ecloud_gosdk.v1/lib/rpc"
	"ecloud_gosdk.v1/lib/url"
)

type UploadConfig struct {
	UpHost    string
	Transport http.RoundTripper
}

type Uploader struct {
	Conn   rpc.Client
	UpHost string
}

func NewUploader(zone int, cfg *UploadConfig) (p Uploader) {

	var uc UploadConfig
	if cfg != nil {
		uc = *cfg
	}

	if uc.UpHost == "" {
		if zone < 0 && zone >= len(conf.Zones) {
			return
		}
		uc.UpHost = conf.Zones[zone].UpHosts[0]
	}

	p.UpHost = uc.UpHost
	p.Conn.Client = &http.Client{Transport: uc.Transport}
	return
}

func NewUploaderWithoutZone(cfg *UploadConfig) (p Uploader) {
	return NewUploader(-1, cfg)
}

// ----------------------------------------------------------

// 根据空间(Bucket)的域名，以及文件的 key，获得 baseUrl。
// 如果空间是 public 的，那么通过 baseUrl 可以直接下载文件内容。
// 如果空间是 private 的，那么需要对 baseUrl 进行私有签名得到一个临时有效的 privateUrl 进行下载。
//
func MakeBaseUrl(domain, key string) (baseUrl string) {
	return "http://" + domain + "/" + url.Escape(key)
}

// ----------------------------------------------------------

// 设置使用这个SDK的应用程序名。userApp 必须满足 [A-Za-z0-9_\ \-\.]*
//
func SetAppName(userApp string) error {

	return conf.SetAppName(userApp)
}

// ----------------------------------------------------------
