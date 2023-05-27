package main

import (
	"fmt"
	"github.com/IskenT/go-rest-api/internal/comment"
	transportHttp "github.com/IskenT/go-rest-api/internal/transport/http"
	"github.com/IskenT/go-rest-api/internal/db"
	log "github.com/sirupsen/logrus"
)

//Run function for creating l (StartUp)
func Run() error{
	fmt.Println("Starting up project")
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil{
		fmt.Println("Failed to migrate database!")
		return err
	}
	fmt.Println("Successfully connected and pinged database!")
	cmtService := comment.NewService(db)
	httpHandler:= transportHttp.NewHandler(cmtService)
	if err:=httpHandler.Serve(); err != nil{
		return err
	}
	return nil
}
func main() {
	fmt.Println("GO REST API")
	if err:=Run(); err!=nil{
		fmt.Println(err)
	}
}