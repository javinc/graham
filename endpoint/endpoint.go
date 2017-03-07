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
		stat := 400

		// format generic error
		if _, ok := err.(*model.Error); !ok {
			err = &model.Error{
				Code:    "GENERIC_ERR",
				Message: err.Error(),
			}
		}

		e := err.(*model.Error)
		// check for internal errors
		if e.Panic {
			stat = 500
		}

		c.JSON(stat, err)

		return
	}

	// okay
	c.JSON(200, data)
}

func parsePayload(c *gin.Context, p interface{}) error {
	m := http.MethodPost + http.MethodPatch + http.MethodDelete
	if strings.Contains(m, c.Request.Method) {
		err := c.BindJSON(p)
		if err != nil {
			err = &model.Error{
				Code:    "ENDPOINT_INVALID_JSON",
				Message: err.Error(),
			}

			print("ERRR")

			return err
		}
	}

	return nil
}
