package middleware

import (
    "net/http"

    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

// RequireAuth redirects to login if no session exists
func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        userID := session.Get("user_id")
        if userID == nil {
            c.Redirect(http.StatusFound, "/login")
            c.Abort()
            return
        }
        // Make user info available to all handlers and templates
        c.Set("user_id", session.Get("user_id"))
        c.Set("user_name", session.Get("user_name"))
        c.Set("role", session.Get("role"))
        c.Set("business_id", session.Get("business_id"))
        c.Set("branch_id", session.Get("branch_id"))
        c.Next()
    }
}

// RequireRole allows access only to specified roles
func RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        for _, r := range roles {
            if r == role {
                c.Next()
                return
            }
        }
        // Authenticated but wrong role — send to their own dashboard
        c.Redirect(http.StatusFound, dashboardFor(role))
        c.Abort()
    }
}

func dashboardFor(role string) string {
    switch role {
    case "director":
        return "/director/dashboard"
    case "manager":
        return "/manager/dashboard"
    default:
        return "/cashier/pos"
    }
}