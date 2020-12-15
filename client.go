package main

import (
	cow11023 "cap-cow11023/proto/cow11023"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)


func main() {
	//创建新的服务
	service := micro.NewService(
		micro.Name("cap.cow11023.client"),
		)
	//初始化
	service.Init()
	capCow11023 :=cow11023.NewCow11023Service(
		"cap.cow11023.server", service.Client())
	str2 := string("跟着水牛走，有肉吃")
	res,err:= capCow11023.AddProduct(context.TODO(),&cow11023.ProductInfo{Message:str2 })
	if err != nil {
		fmt.Println(err)
	}
	 fmt.Println(res)
}
