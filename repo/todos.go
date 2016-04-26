package repo

import (
    _ "github.com/lib/pq"
    "time"
)

var TodoTableSchema = `
CREATE TABLE todos (
    id serial,
    name text,
    completed boolean,
    due timestamp
);`

type Todo struct {
    Id        int64     `json:"id"`
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

func AllTodos() ([]Todo, error) {
    allTodos := "SELECT * FROM todos ORDER BY name ASC"
    todos := []Todo{}
    err := database().Select(&todos, allTodos)
    return todos, err
}

func FindTodo(id int64) (Todo, error) {
    findTodo := "SELECT * FROM todos WHERE id=$1"
    todo := Todo{}
    err = db.Get(&todo, findTodo, id)
    return todo, err
}

func CreateTodo(todo *Todo) error {
    createTodo := "INSERT INTO todos (name, completed, due) VALUES($1, $2, $3) RETURNING id"
    err := db.QueryRow(createTodo, todo.Name, todo.Completed, todo.Due).Scan(&todo.Id)
    return err
}

func DeleteTodo(todoId int64) error {
    deleteTodo := "DELETE FROM todos WHERE id=$1 RETURNING id"
    var id int64
    err := db.QueryRow(deleteTodo, todoId).Scan(&id)
    return err
}

func UpdateTodo(todo *Todo) (Todo, error) {
    updateTodo := "UPDATE todos SET name=$2, completed=$3, due=$4 WHERE id=$1"
    _, err := db.Exec(updateTodo, todo.Id, todo.Name, todo.Completed, todo.Due)
    if err != nil {
        return *todo, err
    }
    return FindTodo(todo.Id)
}
