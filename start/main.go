package main

import (
	"context"
	"entdemo/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/ent?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	// Automatic Migrations​ 自动迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	CreateGraph(context.Background(), client)
	// ctx := context.Background()
	// QueryGithub(ctx, client)
	// QueryArielCars(ctx, client)
	// QueryGroupWithUsers(ctx, client)
}
