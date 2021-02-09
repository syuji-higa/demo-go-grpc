package main

// パッケージを読み込む
import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "local.packages/sample"
)

// 定数を定義する
const (
	address = "localhost:50051"
)

func main() {
	// サーバの接続設定をする
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleClient(conn)

	// サーバにリクエストを送り、レスポンスを表示する

	// リクエストする変数を用意
	var value float64
	// 入力値があるか確認
	if len(os.Args) > 1 {
		// 入力値を string から float64 に変換
		f, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			// float64 に変換できなければ終了
			log.Fatalf("%s\n", err.Error())
		}
		// 入力値をリクエストする変数に代入
		value = f
	} else {
		// 入力値がなければ終了
		log.Fatalf("not found args")
	}
	// サーバに接続する
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 関数にリクエストする
	r, err := c.Square(ctx, &pb.SquareRequest{Value: value})
	if err != nil {
		log.Fatalf("could not sample: %v", err)
	}
	// 入力値を float64 から string 変換
	entry := strconv.FormatFloat(value, 'f', -1, 64)
	// 計算結果を float64 から string 変換
	result := strconv.FormatFloat(r.GetResult(), 'f', -1, 64)
	// 計算結果をコンソールに表示する
	fmt.Printf("%s * %s = %s", entry, entry, result)
}
