package temporal_dto

type DeleteUserWorkflowParam struct {
	UserID   string
	Username string
}

type DeleteUserWorkflowResult struct {
	Success bool
}
