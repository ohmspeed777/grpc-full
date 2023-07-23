package common

type PageInfo struct {
	Page          int64 `json:"page"`
	Size          int64 `json:"size"`
	NumOfPages    int64 `json:"num_of_pages"`
	AllOFEntities int64 `json:"all_of_entities"`
}

type ResponseInfo struct {
	PageInfo PageInfo `json:"page_info"`
	Entities any      `json:"entities"`
}

type Query struct {
	Limit uint    `query:"limit"`
	Page  uint    `query:"page"`
	Query *string `query:"query"`
}
