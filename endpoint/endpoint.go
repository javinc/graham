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
	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/model"
)

func output(c *gin.Context, data interface{}, err error) {
	if err != nil {
		e := err.(*model.Error)
		stat := 400
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
