package client

import (
	"context"

	pb "catdogs.club/back-end/pb"
	"github.com/micro/go-micro"
)

const (
	userClienName = "user.client"
)

var (
	userClient pb.UserService
)

func init() {
	service := micro.NewService(micro.Name(userClienName))
	service.Init()
	userClient = pb.NewUserService("user", service.Client())
}

func Login(req *pb.LoginReq) (*pb.LoginRsp, error) {
	rsp, err := userClient.Login(context.TODO(), req)
	return rsp, err
}

func Regist(req *pb.RegisterReq) (*pb.RegisterRsp, error) {
	rsp, err := userClient.Register(context.TODO(), req)
	return rsp, err
}

func SetProfile(req *pb.SetProfileReq) (*pb.SetProfileRsp, error) {
	rsp, err := userClient.SetProfile(context.TODO(), req)
	return rsp, err
}
