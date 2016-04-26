package repo

import (
    "log"
    "github.com/jmoiron/sqlx"
)

var db, err = sqlx.Connect("postgres", "dbname=postgres sslmode=disable")

func database() *sqlx.DB {
    if db == nil {
        log.Println("WARN: connect(): need to connect!")
        if err != nil {
            log.Fatalln(err)
        }
    }
    return db
}

func InitializeDatabase() {
    db.MustExec(TodoTableSchema)
}
