package grpc

import (
	"desafiogrpc/application/grpc/pb"
	"desafiogrpc/application/grpc/usecase"
	"desafiogrpc/infrastructure/repository"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productRepository := repository.ProductsRepositoryDb{Db: database}
	productsUseCase := usecase.ProductsUseCase{ProductsRepository: productRepository}
	productGrpcService := NewProductGrpcService(productsUseCase)
	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
