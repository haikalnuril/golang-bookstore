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
	Title         *string `json:"title" query:"title"`
	Author        *string `json:"author" query:"author"`
	Genre         *string `json:"genre" query:"genre"`
	PublishedYear *int    `json:"published_year" query:"published_year"`
	Price         *int    `json:"price" query:"price"`
	Page          int     `json:"page" query:"page"`
	Size          int     `json:"size" query:"size"`
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