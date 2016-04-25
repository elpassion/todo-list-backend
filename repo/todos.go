package repo

import (
    "github.com/wojciechko/walletudo-backend/model"
    "fmt"
)

var currentId int

var todos model.Todos

// Give us some seed data
func init() {
    CreateTodo(model.Todo{Name: "Write presentation"})
    CreateTodo(model.Todo{Name: "Host meetup"})
}

func AllTodo() model.Todos {
    return todos
}

func FindTodo(id int) model.Todo {
    for _, t := range todos {
        if t.Id == id {
            return t
        }
    }
    // return empty Todo if not found
    return model.Todo{}
}

func CreateTodo(t model.Todo) model.Todo {
    currentId += 1
    t.Id = currentId
    todos = append(todos, t)
    return t
}

func DestroyTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
