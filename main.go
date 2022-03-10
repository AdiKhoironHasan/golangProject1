package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AdiKhoironHasan/golangProject1/pkg/database"

	integ "github.com/AdiKhoironHasan/golangProject1/internal/integration"
	Repo "github.com/AdiKhoironHasan/golangProject1/internal/repository/postgresql"
	"github.com/AdiKhoironHasan/golangProject1/internal/services"
	handlers "github.com/AdiKhoironHasan/golangProject1/internal/transport/http"
	"github.com/AdiKhoironHasan/golangProject1/internal/transport/http/middleware"

	"github.com/apex/log"
	"github.com/labstack/echo"

	"github.com/spf13/viper"
)

func main() {

	errChan := make(chan error)

	e := echo.New()
	m := middleware.NewMidleware()

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-dev")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	dbhost, dbUser, dbPassword, dbName, dbPort :=
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetString("db.port")

	db, err := database.Initialize(dbhost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatal("Failed to Connect Postgre Database: " + err.Error())
	}

	defer func() {
		err := db.Conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	e.Use(m.CORS)

	sqlrepo := Repo.NewRepo(db.Conn)
	integSrv := integ.NewService()
	srv := services.NewService(sqlrepo, integSrv)
	// srv := services.NewService(sqlrepo)
	handlers.NewHttpHandler(e, srv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- e.Start(":" + viper.GetString("server.port"))
	}()

	e.Logger.Print("Starting ", viper.GetString("appName"))
	err = <-errChan
	log.Error(err.Error())

}
