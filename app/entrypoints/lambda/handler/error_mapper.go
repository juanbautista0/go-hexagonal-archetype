package handler

import (
	"errors"

	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/exceptions"
)

func MapErrorToHTTPResponse(err error) (int, string) {
	var (
		validationError      *exceptions.ValidationError
		businessError        *exceptions.BusinessRuleViolationError
		unauthorizedError    *exceptions.UnauthorizedError
		forbiddenError       *exceptions.ForbiddenError
		notFoundError        *exceptions.NotFoundError
		conflictError        *exceptions.ConflictError
		repositoryError      *exceptions.RepositoryError
		externalServiceError *exceptions.ExternalServiceError
		timeoutError         *exceptions.TimeoutError
	)

	switch {
	case errors.As(err, &validationError):
		return 400, validationError.Error()
	case errors.As(err, &businessError):
		return 422, businessError.Error()
	case errors.As(err, &unauthorizedError):
		return 401, unauthorizedError.Error()
	case errors.As(err, &forbiddenError):
		return 403, forbiddenError.Error()
	case errors.As(err, &notFoundError):
		return 404, notFoundError.Error()
	case errors.As(err, &conflictError):
		return 409, conflictError.Error()
	case errors.As(err, &repositoryError):
		return 500, repositoryError.Error()
	case errors.As(err, &externalServiceError):
		return 502, externalServiceError.Error()
	case errors.As(err, &timeoutError):
		return 504, timeoutError.Error()
	default:
		return 500, "internal server error"
	}
}
