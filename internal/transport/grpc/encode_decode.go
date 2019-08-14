package grpc

import (
	pb "github.com/go-godin/permission-service/api"
	"github.com/go-godin/permission-service/internal/endpoint"
)

func HasResponseEncoder(response endpoint.HasResponse) (*pb.HasResponse, error) {

}

func HasRequestDecoder(request *pb.HasRequest) (endpoint.HasRequest, error) {

}

func GetResponseEncoder(response endpoint.GetResponse) (*pb.GetResponse, error) {

}

func GetRequestDecoder(request *pb.GetRequest) (endpoint.GetRequest, error) {

}

func SetResponseEncoder(response endpoint.SetResponse) (*pb.SetResponse, error) {

}

func SetRequestDecoder(request *pb.HasRequest) (endpoint.HasRequest, error) {

}
