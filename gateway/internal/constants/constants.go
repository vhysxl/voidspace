package constants

const (
	// Errors
	ErrRequestTimeout = "Request timeout"
	ErrInvalidRequest = "Invalid request"
	ErrInternalServer = "Internal server error"
	ErrUnauthorized   = "Unauthorized"

	// Auth
	TokenRefresh    = "Token refreshed successfully"
	LoginSuccess    = "Login successful"
	RegisterSuccess = "Registration successful"
	LogoutSuccess   = "Logout successful"

	// User
	GetProfileSuccess    = "Profile retrieved successfully"
	GetUserSuccess       = "User retrieved successfully"
	UpdateProfileSuccess = "Profile updated successfully"
	DeleteUserSuccess    = "User deleted successfully"
	FollowSuccess        = "User followed successfully"
	UnfollowSuccess      = "User unfollowed successfully"

	// Post
	PostCreated       = "Post created successfully"
	GetPostSuccess    = "Post retrieved successfully"
	UpdatePostSuccess = "Post updated successfully"
	DeletePostSuccess = "Post deleted successfully"

	// Like
	LikeSuccess   = "Post liked successfully"
	UnlikeSuccess = "Post unliked successfully"

	// Comment
	CommentSuccess       = "Comment created successfully"
	GetCommentsSuccess   = "Comments retrieved successfully"
	CommentDeleteSuccess = "Comment deleted successfully"
)
