package handler

import (
	"context"
	"voidspace/users/internal/domain"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (u *UserHandler) UpdateProfile(
	ctx context.Context,
	req *pb.UpdateProfileRequest) (
	*pb.UpdateProfileResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userId, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	updatedProfile := &domain.Profile{
		DisplayName: req.GetDisplayName(),
		Bio:         req.GetBio(),
		AvatarUrl:   req.GetAvatarUrl(),
		BannerUrl:   req.GetBannerUrl(),
		Location:    req.GetLocation(),
	}

	err = u.ProfileUsecase.UpdateProfile(ctx, userId, updatedProfile)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Update Profile")
	}

	return &pb.UpdateProfileResponse{}, nil
}
