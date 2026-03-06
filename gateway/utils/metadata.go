package utils

import "google.golang.org/grpc/metadata"

func MetaDataHandler(userID string, username string) metadata.MD {
	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	return md
}
