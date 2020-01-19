package todo_crud_api_with_go

import "time"

type ToDo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	IsCompleted bool      `json:"is_completed"`
	CreatedOn   time.Time `json:"created_on"`
}
