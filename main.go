package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnka9755/go-03STUDENTAPP/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	router := mux.NewRouter()

	_ = godotenv.Load()

	log := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	_ = db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(log, db)
	userBusiness := user.NewBusiness(log, userRepository)
	userController := user.MakeEndpoints(userBusiness)

	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users", userController.Get).Methods("GET")
	router.HandleFunc("/users", userController.GetAll).Methods("GET")
	router.HandleFunc("/users", userController.Update).Methods("PATCH")
	router.HandleFunc("/users", userController.Delete).Methods("DELETE")

	server := http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*5, "Timeout!"),
		Addr:         "127.0.0.1:5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
