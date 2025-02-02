package handler

import (
	"errors"
	"net/http"

	domain_errors "github.com/c4miloarriagada/keys-be/internal/domain/errors"
	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, err error) {
	var errAuth *domain_errors.UnauthenticatedError

	if errors.As(err, &errAuth) {
		errUnauth := errors.New("unauthenticated")
		writeTypedError(c, http.StatusUnauthorized, errUnauth)
		return
	}

	internalErr := errors.New("internal error")
	writeTypedError(c, http.StatusInternalServerError, internalErr)
}

func writeTypedError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
		"code":  statusCode,
	})
}
