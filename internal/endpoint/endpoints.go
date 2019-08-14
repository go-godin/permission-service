package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"

	"github.com/go-godin/middleware"

	"github.com/go-godin/permission-service/internal/permission"
)

type Set struct {
	SetEndpoint endpoint.Endpoint
	GetEndpoint endpoint.Endpoint
	HasEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc permission.Service, tracer *stdzipkin.Tracer) Set {
	var set endpoint.Endpoint
	{
		set = MakeSetEndpoint(svc)
		if tracer != nil {
			set = middleware.InstrumentZipkin()(set)
			set = zipkin.TraceEndpoint(tracer, "Set")(set)
		}
	}
	var get endpoint.Endpoint
	{
		get = MakeGetEndpoint(svc)
		if tracer != nil {
			get = middleware.InstrumentZipkin()(get)
			get = zipkin.TraceEndpoint(tracer, "Get")(get)
		}
	}
	var has endpoint.Endpoint
	{
		has = MakeHasEndpoint(svc)
		if tracer != nil {
			has = middleware.InstrumentZipkin()(has)
			has = zipkin.TraceEndpoint(tracer, "Has")(has)
		}
	}

	return Set{
		SetEndpoint: set,
		GetEndpoint: get,
		HasEndpoint: has,
	}
}

func MakeSetEndpoint(svc permission.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SetRequest)
		err = svc.Set(ctx, req.Identifier, req.Permission)
		resp := SetResponse{Err: err}
		return resp, nil
	}
}

func MakeGetEndpoint(svc permission.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		perm, err := svc.Get(ctx, req.Identifier)
		resp := GetResponse{
			Permissions: perm,
			Err:         err,
		}
		return resp, nil
	}
}

func MakeHasEndpoint(svc permission.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HasRequest)
		has, err := svc.Has(ctx, req.Identifier, req.Permission)
		resp := HasResponse{
			HasPermission: has,
			Err:           err,
		}
		return resp, nil
	}
}
