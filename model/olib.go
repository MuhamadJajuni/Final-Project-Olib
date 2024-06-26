package model

import "time"

type Book struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	Release_Year int       `json:"release_year"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Borrower struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Address   string    `json:"address"`
}

type Admin struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Transaction struct {
	Id          string    `json:"id"`
	Book_id     string    `json:"book_id"`
	Borrower_id string    `json:"borrower_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Admin_id    string    `json:"admin_id"`
}
