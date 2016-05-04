# todo-list-backend

## Installation
### 1. Download dependencies and compile
```
go get
go build
```
### 2. Start server
##### - Under [localhost:3000](http://localhost:3000) with hot loading `gin`
##### - Under [localhost:3001](http://localhost:3001) without hot loading `./todo-list-backend`

## Usage
- list of todos `GET /todos`
- details of todo `GET /todos/:id`
- create todo `POST /todos`
- remove todo `DELETE /todos/:id`
- update todo `PUT /todos/:id`
