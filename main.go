package main

import (
    "github.com/elpassion/todo-list-backend/web"

    "log"
    "net/http"
)

func main() {
    log.Fatal(http.ListenAndServe(":3001", web.Router()))
}
