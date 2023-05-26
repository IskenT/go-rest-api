package main

import (
	"context"
	"fmt"

	"github.com/IskenT/go-rest-api/internal/comment"
	"github.com/IskenT/go-rest-api/internal/database"
	log "github.com/sirupsen/logrus"
)

//Run function for creating l (StartUp)
func Run() error{
	fmt.Println("Starting up project")
	store, err := database.NewDatabase()
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}
	err = store.MigrateDB()
	if err != nil {
		log.Error("failed to setup database")
		return err
	}
	cmtService := comment.NewService(store)
	fmt.Println(cmtService.GetComment(context.Background(), "b6d5a4f6-bf93-4ad7-8695-f1861d3a0d50"))
	return nil
}
func main() {
	fmt.Println("GO REST API")
	if err:=Run(); err!=nil{
		fmt.Println(err)
	}
}