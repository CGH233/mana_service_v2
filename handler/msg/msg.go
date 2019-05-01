package msg

// GetRequest ...
type GetRequest struct {
	OS   string `form:"os"`
	Page string `form:"page"`
}

// DeleteRequest ...
type DeleteRequest struct {
	OS   string `json:"os"`
	Page string `json:"page"`
}

// CreateResponse ...
type CreateResponse struct{}

// UpdateResponse ...
type UpdateResponse struct{}

// DeleteResponse ...
type DeleteResponse struct{}
