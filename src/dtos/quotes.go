package dtos

type CreateQuotesRequest struct {
	Content string `json:"content" binding:"required"`
	Author  string `json:"author" binding:"required"`
}

type CreateQuotesResponse struct {
	Meta Meta `json:"meta"`
}

type GetQuotesResponse struct {
	Meta Meta        `json:"meta"`
	Data *QuotesData `json:"data"`
}

type QuotesData struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Like    int64  `json:"like"`
}

type LikeQuotesRequest struct {
	ID       int64 `json:"-"`
	Negative bool  `json:"negative"`
}

type LikeQuotesResponse struct {
	Meta Meta            `json:"meta"`
	Data *LikeQuotesData `json:"data"`
}

type LikeQuotesData struct {
	Like int64 `json:"like"`
}
