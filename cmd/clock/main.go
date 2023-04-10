package main

import (
	"context"
	"get-time/internal/rest"
	"get-time/pkg/logger"
	"get-time/pkg/providers/psql"
	"get-time/pkg/service/time_service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	addr  = `:8080`
	pgDSN = getEnv("PG_DSN", "postgresql://postgres:secret@localhost:5433/diplom?sslmode=disable")
)

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	log := logger.GetLogger()

	defer ctxCancel()
	db, err := psql.NewStorage(log, pgDSN)
	if err != nil {
		log.Panic(err)
	}
	if err = db.UpdateSchema(ctx); err != nil {
		log.Infof("migrations error: %v", err)
	}

	service := time_service.NewService(db)
	server := rest.NewServer(log, service, addr)

	server.Run(ctx)

	<-ctx.Done()
	time.Sleep(2 * time.Second)

}

func getEnv(env, defaultValue string) string {
	result := os.Getenv(env)
	if result == "" {
		return defaultValue
	}
	return result
}
