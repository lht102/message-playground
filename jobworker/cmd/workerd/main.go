package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqldrv "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lht102/message-playground/jobworker/config"
	"github.com/lht102/message-playground/jobworker/ent"
	"github.com/lht102/message-playground/jobworker/http/rest"
	"github.com/lht102/message-playground/jobworker/job"
	"github.com/lht102/message-playground/jobworker/nats"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Failed to init logger", err)
	}

	drv, err := sql.Open("mysql", config.GetMySQLDSN(cfg.MySQLCfg))
	if err != nil {
		logger.Fatal("Failed to connect mysql", zap.Error(err))
	}

	db := drv.DB()
	db.SetMaxIdleConns(cfg.MySQLCfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MySQLCfg.MaxOpenConns)
	db.SetConnMaxLifetime(5 * time.Minute)

	driver, err := mysqldrv.WithInstance(
		db,
		&mysqldrv.Config{DatabaseName: cfg.MySQLCfg.Database},
	)
	if err != nil {
		logger.Fatal("Failed to get mysql driver", zap.Error(err))
	}

	path, err := filepath.Abs("./migrations")
	if err != nil {
		logger.Fatal("Failed to get migration path", zap.Error(err))
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+path, "mysql", driver)
	if err != nil {
		logger.Fatal("Failed to init database instance", zap.Error(err))
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal("Failed to migrate", zap.Error(err))
	}

	messageBus, err := nats.NewMessageBus(config.GetNATSURL(cfg.NATSConfig))
	if err != nil {
		logger.Fatal("Failed to connect message bus", zap.Error(err))
	}

	entClient := ent.NewClient(ent.Driver(drv))
	jobService := job.NewService(entClient, messageBus, logger)
	httpServer := rest.NewServer(cfg.RestPort, jobService, logger)

	if err := jobService.RunBackgroundWorker(ctx); err != nil {
		logger.Fatal("Failed to run background worker", zap.Error(err))
	}

	go func() {
		if err := httpServer.Open(); err != nil {
			logger.Fatal("Failed to listen and serve for http server", zap.Error(err))
		}
	}()

	logger.Sugar().Infof("Listening on port %v", cfg.RestPort)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("Shutting down")

	if err := httpServer.Close(); err != nil {
		logger.Fatal("Failed to shutdown http server", zap.Error(err))
	}

	if err := messageBus.Close(); err != nil {
		logger.Fatal("Failed to close message bus connection", zap.Error(err))
	}

	logger.Info("Done")
}
