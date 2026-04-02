package user

import (
	"context"
	"errors"
	temporal_constants "voidspaceGateway/temporal/constants"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

func (s *UserService) DeleteUser(ctx context.Context, userID string, username string) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	param := temporal_dto.DeleteUserWorkflowParam{
		UserID:   userID,
		Username: username,
	}

	run, err := s.TemporalClient.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:        "delete-user-" + userID,
			TaskQueue: s.TemporalService,
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
