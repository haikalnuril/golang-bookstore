package utils

type Paging struct {
	Size        int `json:"size"`
	TotalPage   int `json:"total_page"`
	CurrentPage int `json:"current_page"`
}

type Pageable struct {
	Data   interface{} `json:"data"`
	Paging Paging      `json:"paging"`
}

func PaginateResult(total int, page int, size int) Paging {
	totalPage := (total + size - 1) / size
	return Paging{
		Size:        size,
		TotalPage:   totalPage,
		CurrentPage: page,
	}
}