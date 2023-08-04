package app

import (
	"context"
	"log"
	"one-lab-final/internal/config"
	"one-lab-final/internal/entity"
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

	repo := pgrepo.New(db, cfg)
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

	//Create or update admin user
	admin, err := services.GetUserByCredentials(context.Background(), cfg.ADMIN.Username)
	if err != nil {
		admin = &entity.User{
			Username:  &cfg.ADMIN.Username,
			Email:     &cfg.ADMIN.Email,
			FirstName: &cfg.ADMIN.FirstName,
			LastName:  &cfg.ADMIN.LastName,
			Password: entity.Password{
				Plaintext: &cfg.ADMIN.Password,
			},
			Role: entity.ADMIN,
		}
		log.Print("admin user does not exists, creating new one...")
		services.CreateUser(context.Background(), admin)

	} else {
		log.Print("admin user exists, updating info...")
		services.UpdateUser(context.Background(), admin)
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

	err = services.DeleteUser(context.Background(), admin.ID)
	if err != nil {
		log.Print("error while deleting admin user: ", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
