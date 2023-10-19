package grpc

import (
	"context"
	"desafiogrpc/application/grpc/pb"
	"desafiogrpc/application/grpc/usecase"
	"desafiogrpc/domain/model"
	"hash/fnv"
)

type ProductGrpcService struct {
	ProductUseCase model.ProductsRepositoryInterface
	pb.UnimplementedProductServiceServer
	productUseCase usecase.ProductsUseCase
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := model.NewProduct(in.Name, in.Description, float64(in.Price))
	if err != nil {
		return nil, err
	}

	product, err = p.productUseCase.AddProduct(product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	hash := fnv.New32a()
	hash.Write([]byte(product.ID))
	intID := int32(hash.Sum32())

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          intID,
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.productUseCase.FindAllProducts()
	if err != nil {
		return &pb.FindProductsResponse{}, err
	}

	pbProducts := make([]*pb.Product, len(products))
	for i, product := range products {
		hash := fnv.New32a()
		hash.Write([]byte(product.ID))
		intID := int32(hash.Sum32())

		pbProducts[i] = &pb.Product{
			Id:          intID,
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		}
	}

	return &pb.FindProductsResponse{
		Products: pbProducts,
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductsUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		productUseCase: usecase,
	}
}
