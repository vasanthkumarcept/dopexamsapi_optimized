package start

import (
	"fmt"
	"net/http"
	"recruit/apierrors"

	"github.com/gin-gonic/gin"
)

func HandleResourceNotFoundException(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusNotFound, "Details Not Found", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleBadRequestException(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusBadRequest, "Bad Request", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleInternalServerError(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusInternalServerError, "Server Error", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleUnprocessableEntity(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusUnprocessableEntity, "Malformed JSON request", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleValidationFailed(c *gin.Context, usernames []string) {
	errorDetails := make([]*apierrors.ErrorDetails, len(usernames))
	for i, username := range usernames {
		errorDetails[i] = apierrors.NewErrorDetails(fmt.Sprintf("The Username: %s already exists", username), []string{"The employee ID is already available", "User cannot be inserted"})
	}
	apiError := apierrors.NewAPIError(http.StatusOK, "Validation Error", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleEntityNotFoundException(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusUnprocessableEntity, "Entity Not Found", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleUnauthorizedError(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusUnauthorized, "Unauthorized", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleForbiddenError(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusForbidden, "Forbidden", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleIOException(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusInternalServerError, "I/O Exception", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleBadRequestError(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusBadRequest, "Bad Request", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleMethodNotAllowedError(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusMethodNotAllowed, "Method Not Allowed", errorDetails)
	c.JSON(apiError.Status, apiError)
}

func HandleNotFoundError(c *gin.Context, message string) {
	errorDetails := []*apierrors.ErrorDetails{apierrors.NewErrorDetails(message, nil)}
	apiError := apierrors.NewAPIError(http.StatusNotFound, "Not Found", errorDetails)
	c.JSON(apiError.Status, apiError)
}
