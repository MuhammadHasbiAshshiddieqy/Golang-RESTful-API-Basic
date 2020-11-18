package service

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/gorilla/mux"
	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/MAIN/db"
)

// Body - struct to save request
type Body struct {
	OrderStatus    string         `json:"order_status"`
  ShipmentStatus string         `json:"shipment_status"`
  InvoiceStatus  String         `json:"invoice_status"`
}

func handleRequests(d db.Connection) {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/orderid", getOrderID).Methods("POST")

  return myRouter
}

func unmarshal(r *http.Request) (Body, error) {
	var b Body
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return b, err
	}
	err = json.Unmarshal(rBody, &b)
	if err != nil {
		return b, err
	}
	fmt.Printf("%+v \n", b)
	return b, nil
}
