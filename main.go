package main

import (
    "github.com/wojciechko/walletudo-backend/web"

    "log"
    "net/http"
)

func main() {
    router := web.NewRouter()
    log.Fatal(http.ListenAndServe(":3001", router))
}
