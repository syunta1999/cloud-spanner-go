// main.go
package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"cloud-spanner-go/cmd/di"
)

func main() {
	ctx := context.Background()
	server(ctx)
}

func server(ctx context.Context) {
	app := di.BuildApp()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}
	<-ctx.Done()
	err := app.Stop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// .env ファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
