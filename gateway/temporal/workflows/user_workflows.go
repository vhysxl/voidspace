package workflow

import (
	"time"
	"voidspaceGateway/temporal/activities"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func DeleteUserWorkflow(
	ctx workflow.Context,
	param temporal_dto.DeleteUserWorkflowParam,
) (*temporal_dto.DeleteUserWorkflowResult, error) {

	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    2 * time.Minute,
		ScheduleToStartTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    2 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
			MaximumAttempts:    5,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	actParam := activities.DeleteUserReq{
		UserID:   param.UserID,
		Username: param.Username,
	}

	if err := workflow.ExecuteActivity(
		ctx,
		activities.DeleteUserPostsActivityName,
		actParam,
	).Get(ctx, nil); err != nil {
		return &temporal_dto.DeleteUserWorkflowResult{Success: false},
			temporal.NewApplicationError("DeleteUserPostsActivity failed", "DeleteUserError", err)
	}

	if err := workflow.ExecuteActivity(
		ctx,
		activities.DeleteUserCommentsActivityName,
		actParam,
	).Get(ctx, nil); err != nil {
		return &temporal_dto.DeleteUserWorkflowResult{Success: false},
			temporal.NewApplicationError("DeleteUserCommentsActivity failed", "DeleteUserError", err)
	}

	if err := workflow.ExecuteActivity(
		ctx,
		activities.DeleteUserActivityName,
		actParam,
	).Get(ctx, nil); err != nil {
		return &temporal_dto.DeleteUserWorkflowResult{Success: false},
			temporal.NewApplicationError("DeleteUserActivity failed", "DeleteUserError", err)
	}

	return &temporal_dto.DeleteUserWorkflowResult{Success: true}, nil
}
