package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/ericlagergren/decimal"
	"github.com/gotzmann/grpc/models"
	proto "github.com/gotzmann/grpc/server/gen/proto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/drivers"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	port = ":80"
)

type server struct {
	proto.UnimplementedProductsServer
}

func (s *server) AddProduct(ctx context.Context, in *proto.Product) (*proto.ProductID, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=grpc sslmode=disable")
	if err != nil {
		return nil, err
	}

	var price decimal.Big
	price.SetFloat64(in.Price)

	product := models.Product{
		ProductName: in.Name,
		BrandID: int(in.BrandId),
		CategoryID: int(in.CategoryId),
		ModelYear: int16(in.Year),
		ListPrice: types.Decimal{Big: &price},
	}

	// Blacklist primary key to avoid duplication of IDs
	// Repair nulled sequence counter with SQL if needed:
	// ALTER sequence products_product_id_seq restart with 500;
	err = product.Insert(ctx, db, boil.Blacklist("product_id"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while inserting Product", err)
	}

	log.Printf("\nProduct was added #%v => %v", product.ProductID, product.ProductName)
	return &proto.ProductID{Value: int32(product.ProductID)}, status.New(codes.OK, "").Err()
}

func (s *server) GetProductsByBrand(ctx context.Context, brand *wrapperspb.StringValue) (*proto.ProductsResponse, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=grpc sslmode=disable")
	if err != nil {
		return nil, err
	}
	boil.DebugMode = true

	/* --- Dumb way of selecting with Eager Loading and looping through all collection
	useFiltering := true
	products, err := models.Products(
		Load("Brand"),
	).All(ctx, db) */

	// --- Better way
	useFiltering := false
	products, err := models.Products(
		InnerJoin("brands on brands.brand_id = products.brand_id"),
		Load("Brand", Where("brands.brand_name = ?", brand.Value)),
	).All(ctx, db)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while looking products", err)
	}

	if !useFiltering && len(products) == 0 {
		return nil, status.Errorf(codes.NotFound, "Brand %s has no products", brand.Value)
	}

	brandProducts := make([]*proto.Product, len(products))
	for i, p := range products {

		// Skip the products of other brands for the dumb version
		if useFiltering && brand.Value != p.R.Brand.BrandName {
			continue
		}

		price, _ := p.ListPrice.Float64() // TODO Check errors!
		brandProduct := proto.Product{
			Id: int32(p.ProductID),
			Name: p.ProductName,
			BrandId: int32(p.BrandID),
			CategoryId: int32(p. CategoryID),
			Year: int32(p.ModelYear),
			Price: price,
		}
		brandProducts[i] = &brandProduct
	}

	if len(brandProducts) == 0 {
		return nil, status.Errorf(codes.NotFound, "Brand %s has no products", brand.Value)
	}

	log.Printf("\nReturned %d products for brand %s", len(products), brand.Value)

	return &proto.ProductsResponse{Products: brandProducts}, status.New(codes.OK, "").Err()
}

func main() {
	fmt.Println("[gRPC] Server started...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterProductsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("[gRPC] Server stopped...")
}