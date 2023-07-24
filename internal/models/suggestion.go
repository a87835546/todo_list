package models

import "time"

type Suggestion struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	UserId    int       `json:"user_id"`
	Username  string    `json:"username"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
