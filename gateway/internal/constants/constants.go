package constants

const (
	// Gateway Specific Error

	// Auth
	ErrUsernameRequired = "Username is required"
	TokenRefresh        = "Token refreshed successfully"
	LoginSuccess        = "Login successful"
	RegisterSuccess     = "Registration successful"
	LogoutSuccess       = "Logout successful"

	// User
	ErrNoField             = "No fields to update"
	GetProfileSuccess      = "Profile retrieved successfully"
	GetUserSuccess         = "User retrieved successfully"
	UpdateProfileSuccess   = "Profile updated successfully"
	DeleteUserSuccess      = "User deleted successfully"
	FollowSuccess          = "User followed successfully"
	UnfollowSuccess        = "User unfollowed successfully"
	ListFollowersSuccess   = "Followers retrieved successfully"
	ListFollowingSuccess   = "Following retrieved successfully"

	// Post
	PostCreated        = "Post created successfully"
	GetPostSuccess     = "Post retrieved successfully"
	UpdatePostSuccess  = "Post updated successfully"
	DeletePostSuccess  = "Post deleted successfully"
	GetFeedSuccess     = "Feed retrieved successfully"
	GetUserPostsSuccess  = "User posts retrieved successfully"
	GetLikedPostsSuccess = "Liked posts retrieved successfully"

	// Like
	LikeSuccess   = "Post liked successfully"
	UnlikeSuccess = "Post unliked successfully"

	// Comment
	CommentSuccess       = "Comment created successfully"
	GetCommentsSuccess   = "Comments retrieved successfully"
	CommentDeleteSuccess = "Comment deleted successfully"
)
