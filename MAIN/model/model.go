package model

import (
	"database/sql"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/MAIN/db"
)

type Order struct {
	ID   sql.NullInt64   `json:"order_id"`
}

// Read - read data from database
func Read(ordStat string, shipStat string, invStat string, d db.Connection) (Order, error) {
	o, err := readOrder(orderID, d)
	if err != nil {
		return Order{}, err
	}

  return o, nil
}

func readOrder(ordStat string, shipStat string, invStat string, d db.Connection) (Order, error) {
	o := Order{}
	r := d.Read(orderStmnt(ordStat, shipStat, invStat))

	err := r.Scan(&o.ID)
	if err != nil {
		return o, err
	}

	return o, nil
}

func orderStmnt(ordStat string, shipStat string, invStat string) string {
	tmplt := "SELECT id FROM qa_order_status WHERE status_order = %s AND status_shipment = %s AND status_invoice = %s"
	return fmt.Sprintf(tmplt, ordStat, shipStat, invStat)
}

// NewNullInt64 - create null int64
func NewNullInt64(s int64) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: s,
		Valid: true,
	}
}
