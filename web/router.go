package web

import (
    "github.com/wojciechko/walletudo-backend/web/api"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "log"
)

func Router() http.Handler {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    api.SetApp(router())
    return api.MakeHandler()
}

func router() rest.App {
    router, err := rest.MakeRouter(
        &rest.Route{"GET", "/todos", api.Todos},
        &rest.Route{"GET", "/todos/:todoId", api.FindTodo},
        &rest.Route{"POST", "/todos", api.CreateTodo},
        &rest.Route{"DELETE", "/todos/:todoId", api.DeleteTodo},
        &rest.Route{"PUT", "/todos/:todoId", api.UpdateTodo},
    )
    if err != nil {
        log.Fatalln(err)
    }
    return router
}