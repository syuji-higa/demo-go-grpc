package main

// パッケージを読み込む
import (
	"context"
	"log"
	"math"
	"net"

	pb "local.packages/sample"

	"google.golang.org/grpc"
)

// 定数を定義する
const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSampleServer
}

// 自乗する関数
func (s *server) Square(ctx context.Context, in *pb.SquareRequest) (*pb.SquareResponse, error) {
	val := in.GetValue()
	log.Printf("Received: %v", val)
	result := math.Pow(2.0, val)
	return &pb.SquareResponse{Result: result}, nil
}

func main() {
	// TCP コネクションを開く
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// gRPC サーバを作成する
	s := grpc.NewServer()
	pb.RegisterSampleServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
