/*
  DESIGN RULE
  Package endpoint provides service of /domain from http request

  Endpoint SHOULD ONLY knows internal packages
    [1] /domain

  Endpoint SHOULD ONLY do
    [1] parse input data and pass to /domain
    [2] output formatted results from /domain
*/

package endpoint

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/model"
)

func output(c *gin.Context, data interface{}, err error) {
	if err != nil {
		stat := http.StatusBadRequest

		// format generic error
		if _, ok := err.(*model.Error); !ok {
			err = &model.Error{
				Code:    "GENERIC",
				Message: err.Error(),
			}
		}

		e := err.(*model.Error)
		// check for internal errors
		if e.Panic {
			stat = http.StatusInternalServerError
		}

		c.JSON(stat, err)

		return
	}

	// okay
	c.JSON(http.StatusOK, data)
}

func parsePayload(c *gin.Context, p interface{}) error {
	// check method if needs a payload to read
	m := http.MethodPost + http.MethodPatch + http.MethodDelete
	if strings.Contains(m, c.Request.Method) {
		err := c.BindJSON(p)
		if err != nil {
			// modify error for more info
			err = &model.Error{
				Code:    "ENDPOINT_INVALID_JSON",
				Message: err.Error(),
			}

			return err
		}
	}

	return nil
}
