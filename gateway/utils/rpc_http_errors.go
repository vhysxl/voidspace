package utils

import (
	"net/http"
	"voidspaceGateway/internal/constants"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCErrorToHTTP(err error) (int, string) {
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, constants.ErrInternalServer
	}

	switch st.Code() {
	case codes.NotFound:
		return http.StatusNotFound, st.Message()
	case codes.InvalidArgument:
		return http.StatusBadRequest, st.Message()
	case codes.Unauthenticated:
		return http.StatusUnauthorized, st.Message()
	case codes.PermissionDenied:
		return http.StatusForbidden, st.Message()
	case codes.AlreadyExists:
		return http.StatusConflict, st.Message()
	case codes.Unavailable:
		return http.StatusServiceUnavailable, st.Message()
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, st.Message()
	default:
		return http.StatusInternalServerError, st.Message()
	}
}
