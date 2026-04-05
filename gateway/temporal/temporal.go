package temporal

import (
	"voidspaceGateway/bootstrap"
	"voidspaceGateway/temporal/activities"
	post_activities "voidspaceGateway/temporal/activities/post"
	user_activities "voidspaceGateway/temporal/activities/user"
	workflow "voidspaceGateway/temporal/workflows"
)

func RegisterTemporal(app *bootstrap.Application) {
	userActivities := user_activities.NewUserActivities(
		app.ContextTimeout,
		app.Logger,
		app.UserService.UserClient,
		app.PostService.PostClient,
		app.CommentService.CommentClient,
	)

	postActivities := post_activities.NewPostActivities(
		app.ContextTimeout,
		app.Logger,
		app.PostService.PostClient,
		app.CommentService.CommentClient,
	)

	// registers
	activities.RegisterActivities(app.TemporalService, userActivities, postActivities)
	workflow.RegisterWorkflows(app.TemporalService)
}
