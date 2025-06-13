package response

import (
	"github.com/google/uuid"
)

type TodoResponse struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt string    `json:"created_at"`
}
