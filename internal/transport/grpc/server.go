package grpc

import (
	"context"
	"errors"
	"github.com/go-godin/permission-service/internal/endpoint"
	"github.com/go-kit/kit/transport/grpc"

	pb "github.com/go-godin/permission-service/api"
)

type Server struct {
	SetHandler grpc.Handler
	GetHandler grpc.Handler
	HasHandler grpc.Handler
}

func NewServer(endpoints endpoint.Set, options ...grpc.ServerOption) pb.PermissionServiceServer {
	return &Server{
		SetHandler: grpc.NewServer(
			endpoints.SetEndpoint,
			DecodeSetRequest,
			EncodeSetResponse,
			options...),
		GetHandler: grpc.NewServer(
			endpoints.GetEndpoint,
			DecodeGetRequest,
			EncodeGetResponse,
			options...),
		HasHandler: grpc.NewServer(
			endpoints.HasEndpoint,
			DecodeHasRequest,
			EncodeHasResponse,
			options...),
	}
}

func EncodeHasResponse(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil HasResponse")
	}
	res := response.(endpoint.HasResponse)
	pbResponse, err := HasResponseEncoder(res)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

func DecodeHasRequest(ctx context.Context, pbRequest interface{}) (interface{}, error) {
	if pbRequest == nil {
		return nil, errors.New("nil HasRequest")
	}
	req := pbRequest.(*pb.HasRequest)
	request, err := HasRequestDecoder(req)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeGetResponse(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetResponse")
	}
	res := response.(endpoint.GetResponse)
	pbResponse, err := GetResponseEncoder(res)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil

}

func DecodeGetRequest(ctx context.Context, pbRequest interface{}) (interface{}, error) {
	if pbRequest == nil {
		return nil, errors.New("nil GetRequest")
	}
	req := pbRequest.(*pb.GetRequest)
	request, err := GetRequestDecoder(req)
	if err != nil {
		return nil, err
	}

	return request, nil

}

func EncodeSetResponse(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil SetResponse")
	}
	res := response.(endpoint.SetResponse)
	pbResponse, err := SetResponseEncoder(res)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil

}

func DecodeSetRequest(i context.Context, pbRequest interface{}) (interface{}, error) {
	if pbRequest == nil {
		return nil, errors.New("nil SetRequest")
	}
	req := pbRequest.(*pb.HasRequest)
	request, err := SetRequestDecoder(req)
	if err != nil {
		return nil, err
	}

	return request, nil

}

func (V Server) Set(context.Context, *pb.SetRequest) (*pb.SetResponse, error) {
	panic("implement me")
}

func (V Server) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	panic("implement me")
}

func (V Server) Has(context.Context, *pb.HasRequest) (*pb.HasResponse, error) {
	panic("implement me")
}
