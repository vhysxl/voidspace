package workflow

import (
	"time"
	"voidspaceGateway/temporal/activities/post"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const DeletePostWorkflowName = "DeletePostWorkflow"

func DeletePostWorkflow(ctx workflow.Context, param temporal_dto.DeletePostWorkflowParam) (*temporal_dto.DeletePostWorkflowResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    2,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	actParam := temporal_dto.DeletePostReq(param)
	// 1. Delete Post (Hard Delete)
	errPost := workflow.ExecuteActivity(ctx, post.DeletePostActivity, actParam).Get(ctx, nil)
	if errPost != nil {
		// If post deletion fails, fail the workflow
		return &temporal_dto.DeletePostWorkflowResult{Success: false}, errPost
	}

	// 2. Delete Comments (Eventually Consistent)
	// We wait for it so Temporal can retry it according to the retry policy if it fails
	// Since Post is hard deleted, we don't return an error that would rollback the post,
	// we just let Temporal retry this activity.
	_ = workflow.ExecuteActivity(ctx, post.DeletePostCommentsActivity, actParam).Get(ctx, nil)

	return &temporal_dto.DeletePostWorkflowResult{Success: true}, nil
}
