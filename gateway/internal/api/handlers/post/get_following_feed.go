package post

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	postpb "voidspaceGateway/proto/generated/posts/v1"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *PostHandler) GetFollowingFeed(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("authUser").(*models.AuthUser)

	cursor := c.QueryParam("cursor")
	cursorID := c.QueryParam("cursorid")

	cursorTime, cursorIDInt := utils.ExtractCursor(cursor, cursorID)

	req := &postpb.GetFollowingFeedRequest{}
	if !cursorTime.IsZero() {
		ts := timestamppb.New(cursorTime)
		req.CursorTime = ts
	}
	if cursorIDInt > 0 {
		id := int64(cursorIDInt)
		req.CursorId = &id
	}

	res, err := h.PostService.GetFollowingFeed(ctx, req, user.ID, user.Username)
	if err != nil {
		return utils.HandleDialError(h.Logger, c, err, "failed to fetch following feed")
	}

	return responses.SuccessResponseMessage(c, http.StatusOK, constants.GetFeedSuccess, res)
}
