package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/madeindra/toggl-test/api/v1/deck"
	"github.com/madeindra/toggl-test/internal/config"
	"github.com/madeindra/toggl-test/internal/constant"
	"github.com/madeindra/toggl-test/internal/uuid"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func main() {
	// initialize configurations
	cfg := config.Init()

	// initialize database
	db, err := sql.Open("postgres", cfg.Database.DSN)

	// add logger
	db = sqldblogger.OpenDriver(cfg.Database.DSN, db.Driver(), zerologadapter.New(zerolog.New(os.Stdout)))

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
	uuid := uuid.New()

	// repository and usecase
	deckRepo := deck.NewDeckRepo(db, constant.TableDeck, constant.TableCards)
	deckUsecase := deck.NewDeckUsecase(deckRepo, uuid)

	// router and handler mapping
	deck.NewDeckHandler(router, deckUsecase)

	// initialize server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	// run server
	server.ListenAndServe()
}
