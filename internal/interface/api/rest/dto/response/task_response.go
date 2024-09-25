package response

import "time"

type TaskResponse struct {
	Id          string
	Name        string
	Description string
	StatusId    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ListTaskResponse struct {
	Tasks []*TaskResponse `json:"Tasks"`
}
