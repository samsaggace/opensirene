package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/api/models"
)

// GetHistory is in charge of querying the database to get database history
func (v *ViewsContext) GetHistory(c *gin.Context) {
	var h models.Histories

	if v.GormClient.Find(&h).RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h)
}
