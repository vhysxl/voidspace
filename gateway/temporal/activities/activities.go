package activities

import "voidspaceGateway/bootstrap"

const (
	DeleteUserActivityName         = "DeleteUserActivity"
	DeleteUserPostsActivityName    = "DeleteUserPostsActivity"
	DeleteUserCommentsActivityName = "DeleteUserCommentsActivity"
)

func RegisterActivities(t *bootstrap.TemporalService, ua *UserActivities) {
	t.RegisterActivity(ua.DeleteUserActivity, DeleteUserActivityName)
	t.RegisterActivity(ua.DeleteUserPostsActivity, DeleteUserPostsActivityName)
	t.RegisterActivity(ua.DeleteUserCommentsActivity, DeleteUserCommentsActivityName)
}
