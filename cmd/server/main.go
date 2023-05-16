package main

import "fmt"

//Run function for creating l (StartUp)
func Run() error{
	fmt.Println("Starting up project")
	return nil
}
func main() {
	fmt.Println("GO REST API")
	if err:=Run(); err!=nil{
		fmt.Println(err)
	}
}