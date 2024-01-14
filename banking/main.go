package main

import (
	"banking/handler"
	"banking/logs"
	"banking/repository"
	"banking/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	db := initDatabase()
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepository := repository.NewAccountRepositoryDB(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer)

	router.HandleFunc("/accounts/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
