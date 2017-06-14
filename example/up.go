package main

import (
	"fmt"

	"go-sdk/conf"
	"go-sdk/ecloud"
	"go-sdk/ecloudcli"
)

var (
	// 设置上传到的空间
	bucket = "tantest525"
	key    = "deleteday"
)

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func main() {
	// 初始化AK，SK
	//conf.ACCESS_KEY = "pYKB4NDuaSeq7ziMMBDfcGq7b7-Ueve_1hNFKfQX"
	conf.ACCESS_KEY = "qLkoGKZsHZfZKQZrjwdDWeXiu3S_LsSw"
	//conf.SECRET_KEY = "ndhC36qShV3MNpQa_A5x5qI_6gOEvDJAuviMb7IA"
	conf.SECRET_KEY = "7ra8DkfSSTy5ix49V3dXDZplUyRNj8Ba"
	// 创建一个Client
	c := ecloud.New(0, nil)
	// 设置上传的策略
	policy := &ecloud.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires:         3600,
		InsertOnly:      0,
		DeleteAfterDays: 1,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := ecloudcli.NewUploader(zone, nil)

	var ret PutRet
	// 设置上传文件的路径
	filepath := "/root/work/bin/proxy.toml"
	// 调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFile(nil, &ret, token, "daday121", filepath, nil)
	// 打印返回的信息
	fmt.Println(ret)
	// 打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}
