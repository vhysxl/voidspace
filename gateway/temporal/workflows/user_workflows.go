package workflow

import (
	"strconv"
	"time"
	user_activities "voidspaceGateway/temporal/activities/user"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func DeleteUserWorkflow(
	ctx workflow.Context,
	param temporal_dto.DeleteUserWorkflowParam,
) (*temporal_dto.DeleteUserWorkflowResult, error) {

	userIDInt, err := strconv.Atoi(param.UserID)
	if err != nil {
		return &temporal_dto.DeleteUserWorkflowResult{Success: false}, temporal.NewApplicationError("DeleteUserWorkflow failed", "DeleteUserError", err)
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout:    30 * time.Second,
		ScheduleToStartTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    2,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	actParam := temporal_dto.DeleteUserReq{
		UserID:    param.UserID,
		Username:  param.Username,
		UserIDInt: userIDInt,
	}

	// ── 1. Jalankan semua PARALEL ────────────────────────────
	userFuture := workflow.ExecuteActivity(ctx,
		user_activities.DeleteUserActivity,
		actParam)

	commentsFuture := workflow.ExecuteActivity(ctx,
		user_activities.DeleteUserCommentsActivity,
		actParam)

	postsFuture := workflow.ExecuteActivity(ctx,
		user_activities.DeleteUserPostsActivity,
		actParam)

	// ── 2. Collect semua hasil ───────────────────────────────
	errUser := userFuture.Get(ctx, nil)
	errComments := commentsFuture.Get(ctx, nil)
	errPosts := postsFuture.Get(ctx, nil)

	// ── 3. Kalau semua sukses, selesai ───────────────────────
	if errUser == nil && errComments == nil && errPosts == nil {
		return &temporal_dto.DeleteUserWorkflowResult{Success: true}, nil
	}

	// ── 4. Ada yang gagal → compensate yang SUKSES ──────────
	// Paralel juga, tapi HARUS di-await sebelum return
	var compensateFutures []workflow.Future

	if errUser == nil {
		f := workflow.ExecuteActivity(ctx, user_activities.DeleteUserCompensateActivity, actParam)
		compensateFutures = append(compensateFutures, f)
	}
	if errComments == nil {
		f := workflow.ExecuteActivity(ctx, user_activities.DeleteUserCommentsCompensateActivity, actParam)
		compensateFutures = append(compensateFutures, f)
	}
	if errPosts == nil {
		f := workflow.ExecuteActivity(ctx, user_activities.DeleteUserPostsCompensateActivity, actParam)
		compensateFutures = append(compensateFutures, f)
	}

	// ── 5. Tunggu SEMUA compensate selesai ──────────────────
	// User masih nunggu di UI, jadi tidak boleh fire and forget
	for _, f := range compensateFutures {
		if err := f.Get(ctx, nil); err != nil {
			workflow.GetLogger(ctx).Error("compensate failed", "error", err)
		}
	}

	// ── 6. Return error ke user ──────────────────────────────
	return &temporal_dto.DeleteUserWorkflowResult{Success: false},
		temporal.NewApplicationError("delete account failed, please try again", "DeleteUserError")
}
