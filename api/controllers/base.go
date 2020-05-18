package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/goapi/api/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB 		*gorm.DB
	Router  *mux.Router
}

func (server *Server) Initialize(Dbdriver,DbUser,DbPassword,DbPort,DbHost,DbName string){
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",DbHost,DbPort,DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot Connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)

	}else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(
			&models.User{},
			&models.Post{},
		) //db migration }
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string){
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}