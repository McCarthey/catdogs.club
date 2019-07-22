package client

import (
	pb "catdogs-service/pb"
	"context"

	"github.com/micro/go-micro"
)

const (
	clienName = "post.client"
)

var (
	postClient pb.PostService
)

func init() {
	service := micro.NewService(micro.Name(clienName))
	service.Init()
	postClient = pb.NewPostService("post", service.Client())
}

func SetPost(req *pb.SetPostReq) (*pb.SetPostRsp, error) {
	rsp, err := postClient.Poster(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
