package search

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/service/comment"
	"voidspaceGateway/internal/service/post"
	"voidspaceGateway/internal/service/user"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type SearchHandler struct {
	UserService    *user.UserService
	PostService    *post.PostService
	CommentService *comment.CommentService
	Logger         *zap.Logger
}

func NewSearchHandler(
	userService *user.UserService,
	postService *post.PostService,
	commentService *comment.CommentService,
	logger *zap.Logger,
) *SearchHandler {
	return &SearchHandler{
		UserService:    userService,
		PostService:    postService,
		CommentService: commentService,
		Logger:         logger,
	}
}

func (h *SearchHandler) Search(c echo.Context) error {
	ctx := c.Request().Context()
	query := c.QueryParam("q")
	searchType := c.QueryParam("type")

	if query == "" {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "query parameter 'q' is required")
	}

	switch searchType {
	case "user":
		res, err := h.UserService.SearchUsers(ctx, query)
		if err != nil {
			return utils.HandleDialError(h.Logger, c, err, "failed to search users")
		}
		return responses.SuccessResponseMessage(c, http.StatusOK, "Search users success", res.Users)
	case "post":
		res, err := h.PostService.SearchPosts(ctx, query)
		if err != nil {
			return utils.HandleDialError(h.Logger, c, err, "failed to search posts")
		}
		return responses.SuccessResponseMessage(c, http.StatusOK, "Search posts success", res.Posts)
	case "comment":
		res, err := h.CommentService.SearchComments(ctx, query)
		if err != nil {
			return utils.HandleDialError(h.Logger, c, err, "failed to search comments")
		}
		return responses.SuccessResponseMessage(c, http.StatusOK, "Search comments success", res.Comments)
	default:
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, "invalid search type: must be 'user', 'post', or 'comment'")
	}
}
