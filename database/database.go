package database

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo_app")
    if err != nil {
        log.Fatal(err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal(err)
    }

    log.Println("Asikk database udah konek nih")
}
