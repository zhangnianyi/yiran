package main

import (
	cow11023 "cap-cow11023/proto/cow11023"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

type Cow11023Server struct {}

func (c *Cow11023Server) AddProduct(ctx context.Context, info *cow11023.ProductInfo, product *cow11023.ResponseProduct) error {
	str1 := string("我们的口号是: ")
	product.Answer =  str1+ info.Message
	return nil
}
func main()  {
	//创建新的服务
	service := micro.NewService(
		micro.Name("cap.cow11023.server"),
		)
	//初始化方法
	service.Init()
	//注册服务
	cow11023.RegisterCow11023Handler(service.Server(),new(Cow11023Server))
	//运行服务
	if err := service.Run();err != nil {
		fmt.Println(err)
	}
}
