package main

import (
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
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

	dir, err := atlas.NewLocalDir("./migrations")
	if err != nil {
		// nolint: gocritic
		log.Fatalf("Failed creating atlas migration directory: %v", err)
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
	}

	if err := client.Schema.NamedDiff(ctx, os.Args[1], opts...); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
}
