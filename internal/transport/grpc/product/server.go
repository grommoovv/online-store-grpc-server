package product_grpc

import (
	"context"
	online_storev1 "github.com/grommoovv/online-store-contracts/gen/go/online-store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"online-store-server/internal/service"
)

type serverAPI struct {
	online_storev1.UnimplementedCatalogServer
	product service.Service
}

func Register(grpcServer *grpc.Server, product service.Service) {
	online_storev1.RegisterCatalogServer(grpcServer, &serverAPI{product: product})
}

func (s *serverAPI) GetAll(ctx context.Context, in *online_storev1.GetAllRequest) (*online_storev1.GetAllResponse, error) {
	products, err := s.product.Product.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get all products: %v", err)
	}

	res := make([]*online_storev1.Product, len(products))
	for i, p := range products {
		res[i] = &online_storev1.Product{
			Id:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			Price:       p.Price,
			ImageUrl:    p.ImageURL,
			Category:    p.Category,
		}
	}

	return &online_storev1.GetAllResponse{
		Products: res,
	}, nil
}

func (s *serverAPI) GetById(ctx context.Context, in *online_storev1.GetByIdRequest) (*online_storev1.Product, error) {
	product, err := s.product.Product.GetByID(ctx, int(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get product of id %d: %v", in.Id, err)
	}

	res := &online_storev1.Product{
		Id:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		ImageUrl:    product.ImageURL,
		Category:    product.Category,
	}

	return res, nil
}
