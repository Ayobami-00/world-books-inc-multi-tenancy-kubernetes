package methods

import (
	"time"

	"github.com/google/uuid"
)

func GenerateRandomID() string {

	return uuid.New().String()
}

func GetCurrentTimestamp() int64 {

	currentTime := time.Now()

	// Get the Unix timestamp
	unixTimestamp := currentTime.Unix()

	return unixTimestamp
}
