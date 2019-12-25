package data

type Cursor struct {
	Current          string      `json:"current"`
	EstimatedRecords interface{} `json:"estimatedRecords"`
	HasMore          bool        `json:"hasMore"`
	Next             string      `json:"next"`
	TotalRecords     interface{} `json:"totalRecords"`
}
