package temporal

import (
	"voidspaceGateway/bootstrap"
	"voidspaceGateway/temporal/activities"
	workflow "voidspaceGateway/temporal/workflows"
)

func RegisterTemporal(app *bootstrap.Application) {
	userActivities := activities.NewUserActivities(
		app.ContextTimeout,
		app.Logger,
		app.UserService.UserClient,
		app.PostService.PostClient,
		app.CommentService.CommentClient,
	)

	// registers
	activities.RegisterActivities(app.TemporalService, userActivities)
	workflow.RegisterWorkflows(app.TemporalService)
}
