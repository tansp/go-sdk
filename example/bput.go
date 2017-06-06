package main

import (
	"fmt"

	"ecloud_gosdk.v1/conf"
	"ecloud_gosdk.v1/ecloud"
	"ecloud_gosdk.v1/ecloudcli"
)

var (
	// 设置上传到的空间
	bucket = "demo1111"
	key    = "sdktest15"
)

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func main() {
	// 初始化AK，SK
	conf.ACCESS_KEY = "8bK0SP6WrTrHgGxNbdpWpldUyynqhnqG"
	conf.SECRET_KEY = "Y4GLdWCgjdiy5xUHredUJJVbSZ8FiASH"

	// 创建一个Client
	c := ecloud.New(0, nil)
	// 设置上传的策略
	policy := &ecloud.PutPolicy{
		Scope: bucket + ":" + key,
		//设置Token过期时间
		Expires:    3600,
		InsertOnly: 0,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := ecloudcli.NewUploader(zone, nil)

	var ret PutRet
	// 设置上传文件的路径
	filepath := "/root/work/bin/directory"
	// 调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.RputFile(nil, &ret, token, key, filepath, nil)
	// 打印返回的信息
	fmt.Println(ret)
	// 打印出错信息
	if res != nil {
		fmt.Println("io put failed: ", res)
		return
	}
}
