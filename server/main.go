package main

import (
	"context"
	"log"
	"net"

//	"github.com/gofrs/uuid"
	pb "server/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":80"
)

type server struct {	
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
//	out, err := uuid.NewV4()
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
//	}
//	in.Id = out.String()
//	if s.productMap == nil {
//		s.productMap = make(map[string]*pb.Product)
//	}
//	s.productMap[in.Id] = in
//	log.Printf("Product %v : %v - Added.", in.Id, in.Name)
//	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
	return &pb.ProductID{Value: 1980}, status.New(codes.OK, "").Err()
}

func (s *server) GetProductByBrand(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
//	product, exists := s.productMap[in.Value]
//	if exists && product != nil {
//		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
//		return product, status.New(codes.OK, "").Err()
//	}
//	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}