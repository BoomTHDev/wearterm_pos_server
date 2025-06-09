package custom

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// IsRecordNotFoundError checks if the error is a record not found error
func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// IsDuplicateKeyError checks if the error is a duplicate key error
func IsDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
