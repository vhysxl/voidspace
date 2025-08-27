package models

type GetFeedReq struct {
	UserID   string `json:"userid" validate:"required,min=1"`
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	USerIDs  []int  `json:"userids" validate:"required"`
}

type GetFeedResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"hasmore"`
}
