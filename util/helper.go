package util

import "github.com/satori/go.uuid"

func GenerateUserID() (string, error) {

	userID := uuid.NewV4()
	return userID.String(), nil
}
