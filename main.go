package main

import (
	"currency/api"
	"fmt"
	"os"

	"log"
	"net/http"
	"time"

	"currency/api/currency"
	currencyRepo "currency/pkg/currency/repository"
	currencyUsecase "currency/pkg/currency/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

// Temporary put it here, it should be moved & not inside here
func Handler(DB *sqlx.DB) http.Handler {

	exchangeRepo := currencyRepo.NewMysqlExchangeRateRepository(DB)
	exchangeListRepo := currencyRepo.NewMysqlExchangeListRepository(DB)
	currencyUsecase := currencyUsecase.NewExchangeRateUsecase(exchangeRepo, exchangeListRepo)

	type Routes []api.Route
	var routes = Routes{
		api.Route{
			HandlerFunc: apicurrency.GetCurrencyHandler(currencyUsecase),
			Method:      "GET",
			Path:        "/currency",
			Name:        "GetAllCurrency",
		},
		api.Route{
			HandlerFunc: apicurrency.GetExchangeHandler(currencyUsecase),
			Method:      "GET",
			Path:        "/exchange",
			Name:        "GetExchange",
		},
		api.Route{
			HandlerFunc: apicurrency.PostExchangeHandler(currencyUsecase),
			Method:      "POST",
			Path:        "/exchange",
			Name:        "PostExchange",
		},
		api.Route{
			HandlerFunc: apicurrency.PostExchangeListHandler(currencyUsecase),
			Method:      "POST",
			Path:        "/exchange_list",
			Name:        "PostExchangeList",
		},
		api.Route{
			HandlerFunc: apicurrency.DeleteExchangeListHandler(currencyUsecase),
			Method:      "DELETE",
			Path:        "/exchange_list/{id}",
			Name:        "PostExchangeList",
		},
		api.Route{
			HandlerFunc: apicurrency.GetExchangeListsHandler(currencyUsecase),
			Method:      "GET",
			Path:        "/exchange_list",
			Name:        "GetExchangeList",
		},
	}
	mux := mux.NewRouter()
	for _, route := range routes {
		mux.Methods(route.Method).Name(route.Name).Path(route.Path).Handler(route.HandlerFunc)
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1", "http://localhost", "http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
	})

	return c.Handler(mux)
}

func main() {
	log.Println("Connecting to database ...?")
	db, err := sqlx.Open("mysql", fmt.Sprintf("root:root@tcp(%s)/db_currency?parseTime=true", os.Getenv("DB_URL")))
	if err = db.Ping(); err != nil {
		log.Println("ERROR PING ?", err)
		log.Panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 5)
	handler := Handler(db)
	log.Println("STARTING to Serve")
	// temporary hard code the port, should be moved to the env setting
	log.Fatal(http.ListenAndServe(":8009", handler))
}
