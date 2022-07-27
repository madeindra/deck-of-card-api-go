package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/madeindra/toggl-test/api/v1/deck"
	"github.com/madeindra/toggl-test/internal/config"
	"github.com/madeindra/toggl-test/internal/constant"
)

func main() {
	// initialize configurations
	cfg := config.Init()

	// initialize database
	db, err := sql.Open("postgres", cfg.Database.DSN)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnections)

	// initialize dependencies
	router := mux.NewRouter()
	validator := validator.New()

	// repository and usecase
	deckRepo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	deckUsecase := deck.NewDeckUsecase(deckRepo)

	// router and handler mapping
	deck.NewDeckHandler(router, validator, deckUsecase)

	// initialize server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	// run server
	server.ListenAndServe()
}
