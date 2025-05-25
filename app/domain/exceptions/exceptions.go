package exceptions

import "fmt"

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Msg)
}

type BusinessRuleViolationError struct {
	Msg string
}

func (e *BusinessRuleViolationError) Error() string {
	return fmt.Sprintf("business rule violated: %s", e.Msg)
}

type UnauthorizedError struct {
	Msg string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized: %s", e.Msg)
}

type ForbiddenError struct {
	Msg string
}

func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden: %s", e.Msg)
}

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.Msg)
}

type ConflictError struct {
	Msg string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("conflict: %s", e.Msg)
}

type RepositoryError struct {
	Msg string
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("repository error: %s", e.Msg)
}

type ExternalServiceError struct {
	Msg string
}

func (e *ExternalServiceError) Error() string {
	return fmt.Sprintf("external service error: %s", e.Msg)
}

type TimeoutError struct {
	Msg string
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("timeout: %s", e.Msg)
}
