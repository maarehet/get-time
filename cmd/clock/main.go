package main

import (
	"context"
	"fmt"
	"get-time/api/rest"
	"get-time/config"
	"get-time/providers/psql"
	"get-time/service/time_service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer ctxCancel()
	db, err := psql.NewStorage(config.DbAddress)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.UpdateSchema(ctx); err != nil {
		fmt.Println("migrations error", "->", err)
	}

	service := time_service.NewService(db, db)
	server := rest.NewServer(service)

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
	time.Sleep(2 * time.Second)

}
