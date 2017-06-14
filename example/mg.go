package main

import (
	"fmt"

	"go-sdk/conf"
	"go-sdk/ecloud"
)

var (
	// 设置需要操作的空间
	bucket = "tantest525"
	// 设置需要操作的文件的key
	key = "1111.txt"
	// 设置移动后文件的文件名
	movekey = "efs_server.png_m2"
)

func main() {
	conf.ACCESS_KEY = "qLkoGKZsHZfZKQZrjwdDWeXiu3S_LsSw"
	conf.SECRET_KEY = "7ra8DkfSSTy5ix49V3dXDZplUyRNj8Ba"
	conf.Zones[0].UpHosts = append(conf.Zones[0].UpHosts, "http://192.168.100.98:80")
	// new一个Bucket管理对象
	c := ecloud.New(0, nil)
	p := c.Bucket(bucket)
	// 调用Move方法移动文件
	//	en, res := p.Stat(nil, key)
	entries, _, _, res := p.List(nil, "", "", "", 100)
	// 打印返回值以及出错信息
	fmt.Println("Move success", entries)
	if res == nil {
		fmt.Println("Move success", entries)
	} else {
		fmt.Println("Move failed:", res)
	}
}
