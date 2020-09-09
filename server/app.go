package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	authhttp "github.com/trello-analog/backend/auth/delivery/http"
	postgres2 "github.com/trello-analog/backend/auth/repository/postgres"
	"github.com/trello-analog/backend/auth/usecase"
	"github.com/trello-analog/backend/config"
	"github.com/trello-analog/backend/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	router *mux.Router

	authUseCase *usecase.AuthUseCase
}

func NewApp() *App {
	db := initDB()
	router := initRouter()

	authRepo := postgres2.NewAuthRepository(db)

	return &App{
		authUseCase: usecase.NewAuthUseCase(authRepo),
		router:      router,
	}
}

func (app *App) Run() error {
	cfg := config.GetConfig()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "access-token", "refresh-token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"http://localhost:3000"})

	authRouter := app.router.PathPrefix("/auth").Subrouter()

	authhttp.AuthEndpoints(authRouter, app.authUseCase)

	err := http.ListenAndServe(cfg.Server.Port, handlers.CORS(headers, methods, origin)(app.router))

	if err == nil {
		log.Println("Server successfully launched on " + cfg.Server.Port + " port")
	}

	return err
}

func initDB() *gorm.DB {
	cfg := config.GetConfig()
	dsn := "postgresql://" +
		cfg.Database.User +
		":" +
		cfg.Database.Password +
		"@" +
		cfg.Database.Host +
		":" +
		cfg.Database.Port +
		"/" +
		cfg.Database.Database +
		"?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: services.NewLogger(),
	})

	if err != nil {
		log.Fatal("Database connection was failed!")
	} else {
		log.Println("Database connected successfully!")
	}

	return db
}

func initRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
