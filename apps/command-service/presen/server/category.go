package server

import (
	"context"

	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/application/service"
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen/adapter"
)

type CategoryServer struct {
	adapter adapter.CategoryAdapter
	service service.CategoryService

	// About embedding UnimplementedCategoryCommandServer:
	//
	// Purpose: forward compatibility.
	// 	If a new RPC method is added to the proto in the future,
	// 	this Server doesn't need to implement it to satisfy the interface.
	// 	(The missing method automatically returns an Unimplemented error,
	// 	provided by the embedded type.)
	//
	// Note: embed by value, not by pointer.
	// 	Embedding a pointer risks a nil poointer dereference if it's
	// 	forgotten during initialization. Embedding by value guarantees
	// 	a vlid zero-value struct is always present.
	pbv1.UnimplementedCategoryCommandServer
}

func NewCategoryServer(adapter adapter.CategoryAdapter, service service.CategoryService) pbv1.CategoryCommandServer {
	return &CategoryServer{
		adapter: adapter,
		service: service,
	}
}

func (ins *CategoryServer) Create(ctx context.Context, param *pbv1.CategoryUpParam) (*pbv1.CategoryUpResult, error) {
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Add(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(category), nil
	}
}

func (ins *CategoryServer) Update(ctx context.Context, param *pbv1.CategoryUpParam) (*pbv1.CategoryUpResult, error) {
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Update(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(category), nil
	}
}

func (ins *CategoryServer) Delete(ctx context.Context, param *pbv1.CategoryUpParam) (*pbv1.CategoryUpResult, error) {
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Delete(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(category), nil
	}
}
