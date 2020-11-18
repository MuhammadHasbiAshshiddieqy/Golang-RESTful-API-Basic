package service

import (
	"fmt"
  "log"
	"net/http"
  "encoding/json"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/MAIN/db"
	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/MAIN/model"
)

func getOrderID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint Hit: getOrderID")
  b, err := unmarshal(r)
	if err != nil {
		fmt.Println(err)
	}

	o, err := model.Read(b.OrderStatus, b.ShipmentStatus, b.InvoiceStatus , d)
	if err != nil {
		log.Printf("Error read model: %s", err)
	}

  json.NewEncoder(w).Encode(o)
}
