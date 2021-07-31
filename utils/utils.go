package utils

import (
	uuid "github.com/satori/go.uuid"
)

func GenUuid() string {
	return uuid.NewV4().String()
}
