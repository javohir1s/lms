package main

import (
	"context"
	"fmt"
	"lms_back/api"
	"lms_back/config"
	"lms_back/service"
	"lms_back/storage/postgres"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store)

	c := api.New(store, services)

	fmt.Println("programm is running on localhost:8080...")
	c.Run(":8080")
}
