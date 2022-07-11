package main

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lht102/message-playground/jobworker/config"
	"github.com/lht102/message-playground/jobworker/ent"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	client, err := ent.Open("mysql", config.GetMySQLDSN(cfg.MySQLCfg))
	if err != nil {
		log.Fatalf("Failed connecting to mysql: %v", err)
	}
	defer client.Close()

	dir, err := migrate.NewLocalDir("../../migrations")
	if err != nil {
		// nolint: gocritic
		log.Fatalf("Failed creating atlas migration directory: %v", err)
	}

	if err := client.Schema.NamedDiff(ctx, os.Args[1], schema.WithDir(dir), schema.WithSumFile()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
}
