package model

import (
	"context"
	"fmt"
	hello "grpc2/server/proto"
)

type Say struct {

}

func(s *Say)Hello(ctx context.Context,req *hello.Req,res *hello.Res)error{
	res.Name = "hello," + req.Name
	fmt.Println("req:",req,"res:",res)
	return nil
}