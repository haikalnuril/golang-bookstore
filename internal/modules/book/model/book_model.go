package model

type BookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"published_year"`
	Price         int    `json:"price"`
}

type UpdateBookRequest struct {
	ID            string  `json:"id"`
	Title         *string `json:"title"`
	Author        *string `json:"author"`
	Genre         *string `json:"genre"`
	PublishedYear *int    `json:"published_year"`
	Price         *int    `json:"price"`
}

type SearchBookRequest struct {
	Title         *string `json:"title,omitempty"`
	Author        *string `json:"author,omitempty"`
	Genre         *string `json:"genre,omitempty"`
	PublishedYear *int    `json:"published_year,omitempty"`
	Price         *int    `json:"price,omitempty"`
	Page          *int    `json:"page,omitempty"`
	Size          *int    `json:"size,omitempty"`
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