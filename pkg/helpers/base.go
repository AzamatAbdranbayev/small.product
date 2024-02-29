package helpers

import "github.com/google/uuid"

func CheckValidUuid(id string) error {
	_, err := uuid.Parse(id)
	return err
}
