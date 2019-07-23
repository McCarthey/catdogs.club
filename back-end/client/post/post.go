package client

import (
	"context"

	pb "catdogs.club/back-end/pb"
	"github.com/micro/go-micro"
)

const (
	postClienName = "post.client"
)

var (
	postClient pb.PostService
)

func init() {
	service := micro.NewService(micro.Name(postClienName))
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
