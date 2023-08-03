package app

import (
	"context"
	"log"
	"one-lab-final/internal/config"
	"one-lab-final/internal/handler"
	"one-lab-final/internal/repository/pgrepo"
	"one-lab-final/internal/service"
	"one-lab-final/pkg/httpserver"
	"one-lab-final/pkg/store/postgres"
	"os"
	"os/signal"

	"github.com/procyon-projects/chrono"
)

func Run(cfg *config.Config) error {
	db, err := postgres.ConnectDB(
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithUsername(cfg.DB.Username),
		postgres.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}

	log.Println("connection success")

	repo := pgrepo.New(db)
	services := service.New(repo, cfg)
	handler := handler.New(services, cfg)
	server := httpserver.New(
		handler.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	//Delete expired tokens at 00:00 on Sunday
	taskScheduler := chrono.NewDefaultTaskScheduler()
	_, err = taskScheduler.ScheduleWithCron(func(ctx context.Context) {
		services.DeleteExpiredTokens(ctx)
		log.Println("expired tokens are deleted")
	}, "0 0 0 * * 0")
	if err != nil {
		log.Printf("scheduling task error: %s", err.Error())
		return err
	}

	//Refresh rating of books every 3 hours
	_, err = taskScheduler.ScheduleWithCron(func(ctx context.Context) {
		services.RefreshBooksRating(ctx)
		log.Println("ratings of books are refreshed")
	}, "0 0 */3 * * *")
	if err != nil {
		log.Printf("scheduling task error: %s", err.Error())
		return err
	}

	server.Start()
	log.Println("server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
