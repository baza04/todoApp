package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	todoapp "github.com/baza04/todoApp"
	"github.com/baza04/todoApp/pkg/handler"
	"github.com/baza04/todoApp/pkg/repository"
	"github.com/baza04/todoApp/pkg/service"
)

// @title Todo App API
// @version 1.0
// @description API Server for TODO Application

// @host localhost:8001
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

type config struct {
	port string
	DB
}

type DB struct {
	Username string
	password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func main() {
	var (
		db *sqlx.DB
		// conf *config
		err error
	)

	logrus.SetFormatter(new(logrus.JSONFormatter))

	// conf, err := initConfig()
	// if err != nil {
	// 	logrus.Fatalf("error initializating configs: %s", err.Error())
	// }
	// logrus.Fatalf("conf: %#v\n", conf)

	// if conf, err = initConfig(); err != nil {
	// 	logrus.Fatalf("error initializating configs: %v", err)
	// }

	if err = godotenv.Load("./configs/config.yml"); err != nil {
		logrus.Errorf("error loading env variables: %v", err)
	}

	// if value, ok := os.LookupEnv("DB_PASSWORD"); !ok || value == "" {
	// 	logrus.Fatalf("error cannot initializing database password")
	// }

	// or use value from config struct -> Host:     conf.DB.Host,
	if db, err = repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "postgres",
		// SSLMode:  false,
		// Host:     viper.GetString("db.host"),
		// Port:     viper.GetString("db.port"),
		// Username: viper.GetString("db.username"),
		// DBName:   viper.GetString("db.dbname"),
		// SSLMode:  viper.GetString("db.sslmode"),
		// Password: os.Getenv("DB_PASSWORD"),
	}); err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	srv := new(todoapp.Server)
	go func() {
		if err = srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			// if err = srv.Run(conf.port, handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error on db connection close: %s", err.Error())
	}
}

// func initConfig() (*config, error) {
func initConfig() error {
	// viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	return fmt.Errorf("config init with viper failed: %w", viper.ReadInConfig())

	// conf := new(config)
	// if err := viper.Unmarshal(&conf); err != nil {
	// 	return nil, err
	// }

	// if err := godotenv.Load(); err != nil {
	// 	return nil, err
	// }

	// if value, ok := os.LookupEnv("DB_PASSWORD"); !ok || value == "" {
	// 	return nil, errors.New("error cannot initializing database password")
	// } else {
	// 	conf.DB.password = value
	// }

	// return conf, nil
}
