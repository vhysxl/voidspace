package activities

import (
	"voidspaceGateway/bootstrap"
	"voidspaceGateway/temporal/activities/post"
	"voidspaceGateway/temporal/activities/user"
)

func RegisterActivities(t *bootstrap.TemporalService, ua *user.UserActivities, pa *post.PostActivities) {
	t.RegisterActivity(ua.DeleteUserActivity, user.DeleteUserActivity)
	t.RegisterActivity(ua.DeleteUserPostsActivity, user.DeleteUserPostsActivity)
	t.RegisterActivity(ua.DeleteUserCommentsActivity, user.DeleteUserCommentsActivity)

	// Compensate Activities
	t.RegisterActivity(ua.DeleteUserCompensateActivity, user.DeleteUserCompensateActivity)
	t.RegisterActivity(ua.DeleteUserCommentsCompensateActivity, user.DeleteUserCommentsCompensateActivity)
	t.RegisterActivity(ua.DeleteUserPostsCompensateActivity, user.DeleteUserPostsCompensateActivity)

	// Post Activities
	t.RegisterActivity(pa.DeletePostActivity, post.DeletePostActivity)
	t.RegisterActivity(pa.DeletePostCommentsActivity, post.DeletePostCommentsActivity)
}
