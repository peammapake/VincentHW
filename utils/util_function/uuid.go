package util_function

import "github.com/google/uuid"

func PointerUUIDToPointerString(input *uuid.UUID) *string {
	if input == nil {
		return nil
	} else {
		output := input.String()
		return &output
	}
}
