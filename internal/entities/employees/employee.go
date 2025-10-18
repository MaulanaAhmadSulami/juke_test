package employeeEntity

import "time"

type Employee struct {
	ID         int64   `json:"id" example:"1"`
	Name       string  `json:"name" example:"John Doe"`
	Email      string  `json:"email" example:"john.doe@example.com"`
	Position   string  `json:"position" example:"Software Engineer"`
	Salary     float64 `json:"salary" example:"100000"`
	CreatedAt time.Time `json:"created_at"`
}