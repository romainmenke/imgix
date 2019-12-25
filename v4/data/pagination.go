package data

type Pagination struct {
	CurrentPage     int  `json:"currentPage"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	NextPage        int  `json:"nextPage"`
	PageSize        int  `json:"pageSize"`
	PreviousPage    int  `json:"previousPage"`
	TotalPages      int  `json:"totalPages"`
	TotalRecords    int  `json:"totalRecords"`
}
