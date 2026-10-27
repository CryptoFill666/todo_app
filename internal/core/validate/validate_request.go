package validate

import "fmt"

func ValidateCreateTaskRequest(title string, description *string) error {
	if title == "" {
		return fmt.Errorf("title is required")
	}
	if description != nil && *description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	return nil
}
