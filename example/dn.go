package main

import (
	"ecloud_gosdk.v1/conf"
	"ecloud_gosdk.v1/ecloud"
)

var (
	// 指定私有空间的域名
	domain = "aaaah3.bucket.gosuncloud.com"
	// 指定文件的key
	key = "11.jpg"
)

// 调用封装好的downloadUrl方法生成一个下载链接
func downloadUrl(domain, key string) string {
	// 调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := ecloud.MakeBaseUrl(domain, key)
	policy := ecloud.GetPolicy{}
	// 生成一个client对象
	c := ecloud.New(0, nil)
	// 调用MakePrivateUrl方法返回url
	return c.MakePrivateUrl(baseUrl, &policy)
}
func main() {
	conf.ACCESS_KEY = "FhUXs-GGwIqx149mSMId1e4fMzH9quwN"
	conf.SECRET_KEY = "LOuP6dH8CMNVeeLvfeNK6pOQzYFPkoC2"
	// 打印出下载链接
	println(downloadUrl(domain, key))
}
