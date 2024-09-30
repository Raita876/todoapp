package response

import "time"

type TaskResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StatusId    int       `json:"status_id"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

type ListTaskResponse struct {
	Tasks []*TaskResponse `json:"tasks"`
}
