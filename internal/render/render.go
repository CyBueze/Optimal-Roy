package render

import (
    "net/http"

    "github.com/a-h/templ"
    "github.com/gin-gonic/gin"
)

func Page(c *gin.Context, status int, component templ.Component) {
    c.Status(status)
    if err := component.Render(c.Request.Context(), c.Writer); err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
    }
}