package main

import (
    "github.com/wojciechko/walletudo-backend/web"

    "log"
    "net/http"
)

func main() {
    log.Fatal(http.ListenAndServe(":3001", web.Router()))
}
