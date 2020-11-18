package main

import (
	"os"
	"fmt"
	"log"
  "errors"
	"net/http"
  "encoding/json"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/MAIN/db"
  "github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/MAIN/service"
)

func main() {
	err := checkEnv()
	if err != nil {
		log.Fatal(err)
	}

	d := db.Connect()
	defer d.Close()

  router := service.handleRequests(d)

	fmt.Printf("Starting server at port 8080\n")
  log.Fatal(http.ListenAndServe(":8080", router))
}

func checkEnv() error {
	log.Println("Checking environment.....")

	_, ok := os.LookupEnv("DATA_SOURCE")
	if !ok {
		return errors.New("env:data_source does not exist")
	}

	_, ok = os.LookupEnv("APIGATEWAY_URL")
	if !ok {
		return errors.New("env:cred_url does not exist")
	}

	log.Println("Environment is complete.....")
	return nil
}
