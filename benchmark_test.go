package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/taako-502/go-mongodb-benchmark/pkg/benchmark"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client
var benchmarkPattern = []int{100, 1000, 10000, 100000}

func init() {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

// BenchmarkA Aのベンチマークを取る
func BenchmarkA(b *testing.B) {
	ctx := context.Background()
	// コレクションを初期化
	con := client.Database("go_mongodb_benchmark_template").Collection("sample")
	benchmark.Cleanup(ctx, con)

	for _, n := range benchmarkPattern {
		b.Run("Benchmark_"+fmt.Sprint(n), func(b *testing.B) {
			// TODO: 前処理

			b.ResetTimer()
			for b.Loop() {
				// TODO: ベンチマーク対象の処理
			}
		})
	}
}
