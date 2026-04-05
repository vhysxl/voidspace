package post

import (
	"context"

	"errors"
	"fmt"
	temporal_constants "voidspaceGateway/temporal/constants"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ps *PostService) Delete(ctx context.Context, postID int64, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	param := temporal_dto.DeletePostWorkflowParam{
		PostID:   postID,
		Username: username,
		UserID:   userID,
	}

	workflowID := fmt.Sprintf("delete-post-%d", postID)

	run, err := ps.TemporalClient.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:                       workflowID,
			TaskQueue:                ps.TemporalService,
			WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
			WorkflowIDConflictPolicy: enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING,
		},
		temporal_constants.DeletePostWorkflowName,
		param,
	)
	if err != nil {
		ps.Logger.Error("failed to execute workflow", zap.Error(err))
		return err
	}

	var res temporal_dto.DeletePostWorkflowResult
	if err := run.Get(ctx, &res); err != nil {
		var appErr *temporal.ApplicationError
		if errors.As(err, &appErr) {
			switch appErr.Type() {
			case "PermissionDenied":
				return status.Error(codes.PermissionDenied, appErr.Message())
			case "NotFound":
				return status.Error(codes.NotFound, appErr.Message())
			default:
				return status.Error(codes.Internal, appErr.Message())
			}
		}
		ps.Logger.Error("workflow failed", zap.Error(err))
		return err
	}

	if !res.Success {
		return errors.New("delete post workflow failed")
	}

	return nil
}
