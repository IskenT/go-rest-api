package main

import (
	"context"
	"fmt"
	"github.com/IskenT/go-rest-api/internal/comment"
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
	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID: "b6d5a4f6-bf93-4ad7-8695-f1861d3a0d50",
			Slug: "maual-test",
			Author: "Elliot",
			Body: "Hello World!",
		},
	)
	fmt.Println(cmtService.GetComment(
		context.Background(), 
		"b6d5a4f6-bf93-4ad7-8695-f1861d3a0d50"))
	return nil
}
func main() {
	fmt.Println("GO REST API")
	if err:=Run(); err!=nil{
		fmt.Println(err)
	}
}