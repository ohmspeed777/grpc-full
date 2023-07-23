package domain

import "math"



type Query struct {
	Limit uint `json:"limit" path:"limit" query:"limit"`
	Page  uint `json:"page" path:"page" query:"page"`
}

func (q *Query) GetLimit() int64 {
	if q.Limit == 0 {
		q.Limit = 20
	}

	return int64(q.Limit)
}

func (q *Query) GetPage() int64 {
	if q.Page == 0 {
		q.Page = 1
	}

	return int64(q.Page)
}

func (q *Query) GetSkip() int64 {
	if q.Limit == 0 {
		q.Limit = 20
	}

	if q.Page == 0 {
		q.Page = 1
	}

	return int64((q.Page - 1) * q.Limit)
}

type GetOneReq struct {
	ID string `json:"id" path:"id" query:"id" param:"id"`
}

type PageInfo struct {
	Page          int64 `json:"page"`
	Size          int64 `json:"size"`
	NumOfPages    int64 `json:"num_of_pages"`
	AllOFEntities int64 `json:"all_of_entities"`
}

type ResponseInfo[T any] struct {
	PageInfo PageInfo `json:"page_info"`
	Entities T        `json:"entities"`
}

func NewPagination[T any](entities T, page, size, counter int64) ResponseInfo[T] {
	res := ResponseInfo[T]{}
	res.Entities = entities
	res.PageInfo.NumOfPages = int64(math.Ceil(float64(counter) / float64(size)))
	res.PageInfo.Page = page
	res.PageInfo.Size = size
	res.PageInfo.AllOFEntities = counter
	return res
}
