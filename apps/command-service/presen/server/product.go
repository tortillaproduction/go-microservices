package server

import (
	"context"

	pbv1 "github.com/tortillaproduction/go-microservices/api/pb/v1"
	"github.com/tortillaproduction/go-microservices/apps/command-service/application/service"
	"github.com/tortillaproduction/go-microservices/apps/command-service/presen/adapter"
)

type ProductServer struct {
	adapter adapter.ProductAdapter
	service service.ProductService

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
	pbv1.UnimplementedProductCommandServer
}

func NewProductServer(adapter adapter.ProductAdapter, service service.ProductService) pbv1.ProductCommandServer {
	return &ProductServer{
		adapter: adapter,
		service: service,
	}
}

func (ins *ProductServer) Create(ctx context.Context, param *pbv1.ProductUpParam) (*pbv1.ProductUpResult, error) {
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Add(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(product), nil
	}
}

func (ins *ProductServer) Update(ctx context.Context, param *pbv1.ProductUpParam) (*pbv1.ProductUpResult, error) {
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Update(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(product), nil
	}
}

func (ins *ProductServer) Delete(ctx context.Context, param *pbv1.ProductUpParam) (*pbv1.ProductUpResult, error) {
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Delete(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(product), nil
	}
}
