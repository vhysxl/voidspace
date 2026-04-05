package user

import (
	"context"
	"errors"
	"strconv"
	temporal_constants "voidspaceGateway/temporal/constants"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

func (s *UserService) DeleteUser(ctx context.Context, userID string, username string) error {
	user, err := s.GetCurrentUser(ctx, userID, username)
	if err != nil {
		s.Logger.Error("failed to get user", zap.Error(err))
		return err
	}

	param := temporal_dto.DeleteUserWorkflowParam{
		UserID:   strconv.Itoa(user.ID),
		Username: username,
	}

	run, err := s.TemporalClient.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:                       "delete-user-" + userID,
			TaskQueue:                s.TemporalService,
			WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
			WorkflowIDConflictPolicy: enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING,
		},
		temporal_constants.DeleteUserWorkflowName,
		param,
	)

	if err != nil {
		s.Logger.Error("failed to execute workflow", zap.Error(err))
		return err
	}

	var res temporal_dto.DeleteUserWorkflowResult
	if err := run.Get(ctx, &res); err != nil {
		s.Logger.Error("workflow failed", zap.Error(err))
		return err
	}

	if !res.Success {
		return errors.New("delete user workflow failed")
	}

	return nil
}
