package temporal_dto

type DeleteUserWorkflowParam struct {
	UserID   string
	Username string
}

type DeleteUserWorkflowResult struct {
	Success bool
}

// ===================================== Delete User DTOs =====================================
type DeleteUserReq struct {
	UserIDInt int
	UserID    string
	Username  string
}

// ===================================== Delete Post DTOs =====================================
type DeletePostWorkflowParam struct {
	PostID   int64
	Username string
	UserID   string
}

type DeletePostWorkflowResult struct {
	Success bool
}

type DeletePostReq struct {
	PostID   int64
	Username string
	UserID   string
}
