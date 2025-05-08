package model

type BookRequest struct {
	Title         string
	Author        string
	Genre         string
	PublishedYear int
}

type BookResponse struct {
	ID            string
	Title         string
	Author        string
	Genre         string
	PublishedYear int
	CreatedAt	 string
	UpdatedAt	 string
}