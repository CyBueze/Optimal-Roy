package handlers

import (
	"net/http"
	"github.com/CyBueze/org/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, t templ.Component) {
	c.Status(status)
	if err := t.Render(c.Request.Context(), c.Writer); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func Index(c *gin.Context) {
	render(c, http.StatusOK, views.Index("org"))
}