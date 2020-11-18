package main

import (
    "os"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")
    // Access format "username:password@tcp(127.0.0.1:3306)/dbname"
    db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE"))
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

}
