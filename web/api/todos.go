package api

import (
    "github.com/wojciechko/walletudo-backend/repo"
    "github.com/ant0ine/go-json-rest/rest"
    "strconv"
    "net/http"
)

func Todos(writer rest.ResponseWriter, request *rest.Request) {
    todos, err := repo.AllTodos()
    if err != nil {
        rest.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := writer.WriteJson(todos); err != nil {
        rest.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func FindTodo(writer rest.ResponseWriter, request *rest.Request) {
    todoId, err := strconv.ParseInt(request.PathParam("todoId"), 10, 64)
    if err != nil {
        rest.Error(writer, "Invalid :todoId path param", http.StatusBadRequest)
        return
    }
    todo, err := repo.FindTodo(todoId)
    if err != nil {
        rest.NotFound(writer, request)
        return
    }
    if err := writer.WriteJson(todo); err != nil {
        rest.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func CreateTodo(writer rest.ResponseWriter, request *rest.Request) {
    todo := repo.Todo{}
    err := request.DecodeJsonPayload(&todo)
    if err != nil {
        rest.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    if len(todo.Name) == 0 {
        rest.Error(writer, "Todo name required", http.StatusBadRequest)
        return
    }
    err = repo.CreateTodo(&todo)
    if err != nil {
        rest.Error(writer, err.Error(), http.StatusInternalServerError)
    }
    writer.WriteJson(todo)
}

func DeleteTodo(writer rest.ResponseWriter, request *rest.Request) {
    todoId, err := strconv.ParseInt(request.PathParam("todoId"), 10, 64)
    if err != nil {
        rest.Error(writer, "Invalid :todoId path param", http.StatusBadRequest)
        return
    }
    err = repo.DeleteTodo(todoId)
    if err != nil {
        rest.NotFound(writer, request)
        return
    }
    writer.WriteHeader(http.StatusOK)
}
