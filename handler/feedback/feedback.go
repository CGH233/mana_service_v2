package feedback

import "github.com/asynccnu/mana_service_v2/model"

type CreateResponse struct{}

type ListRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}
type ListResponse struct {
	TotalCount int64                 `json:"totalCount"`
	List       []*model.FeedbackItem `json:"list"`
}
