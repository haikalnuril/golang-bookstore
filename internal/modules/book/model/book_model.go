package model

type BookRequest struct {
	Title         string
	Author        string
	Genre         string
	PublishedYear int
	Price         int
}

type BookResponse struct {
	ID            string
	Title         string
	Author        string
	Genre         string
	PublishedYear int
	Price         int
	CreatedAt     string
	UpdatedAt     string
}
