package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/javinc/graham/model"
)

// Output response
func Output(c *gin.Context, data interface{}, err error) {
	if err != nil {
		OutputError(c, err)

		return
	}

	// okay
	c.JSON(http.StatusOK, data)
}

// OutputError response
func OutputError(c *gin.Context, err error) {
	stat := http.StatusBadRequest

	// format generic error
	if _, ok := err.(*model.Error); !ok {
		err = &model.Error{
			Name:    "GENERIC",
			Message: err.Error(),
		}
	}

	e := err.(*model.Error)
	// check for internal errors
	if e.Panic {
		stat = http.StatusInternalServerError
	}

	c.JSON(stat, err)
}

// OutputNotFound response with 404
func OutputNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, &model.Error{
		Name:    "NOT_FOUND",
		Message: "404 not found",
	})
}

// ParsePayload parse payload with error handling
func ParsePayload(c *gin.Context, p interface{}) error {
	// check method if needs a payload to read
	m := http.MethodPost + http.MethodPatch
	if strings.Contains(m, c.Request.Method) {
		err := c.BindJSON(p)
		if err != nil {
			// modify error for more info
			err = &model.Error{
				Name:    "INVALID_JSON",
				Message: err.Error(),
			}

			return err
		}
	}

	return nil
}
