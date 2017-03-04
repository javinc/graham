package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/graham/domain"
)

// FindUser endpoint handler
func FindUser(c *gin.Context) {
	d, err := domain.FindUser(c)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, d)
}

// GetUser endpoint handler
func GetUser(c *gin.Context) {
	id := c.Param("id")
	d, _ := domain.GetUser(c, id)

	c.JSON(200, d)
}