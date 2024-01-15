package main

import (
	"banking/handler"
	"banking/logs"
	"banking/repository"
	"banking/services"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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
	customerHandler := handler.NewCustomerHandlerFiber(customerService)

	accountRepository := repository.NewAccountRepositoryDB(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandlerFiber := handler.NewAccountHandlerFiber(accountService)

	// router := mux.NewRouter()
	app := fiber.New()

	app.Get("/accounts/:customer_id/accounts", accountHandlerFiber.GetAccounts)
	app.Post("/accounts/:customer_id/accounts", accountHandlerFiber.NewAccount)

	app.Get("/customers", customerHandler.GetCustomers)
	app.Get("/customers/:customer_id", customerHandler.GetCustomer)

	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	// http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
	app.Listen(":8000")
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
