package main

import (
	"fmt"
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
	return nil
}
func main() {
	fmt.Println("GO REST API")
	if err:=Run(); err!=nil{
		fmt.Println(err)
	}
}