package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/xXHachimanXx/Imersao-FullCycle-2.0/application/route"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[1])
}
