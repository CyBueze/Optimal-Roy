package authhandler

import (
    "net/http"

    "github.com/CyBueze/org/internal/models"
    "github.com/CyBueze/org/internal/render"
    authpages "github.com/CyBueze/org/views/pages/auth"

    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    csrf "github.com/utrack/gin-csrf"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func ShowLogin(c *gin.Context) {
    session := sessions.Default(c)
    if session.Get("user_id") != nil {
        role := session.Get("role").(string)
        c.Redirect(http.StatusFound, dashboardFor(role))
        return
    }

    errorMsg := ""
    if c.Query("error") == "invalid" {
        errorMsg = "Invalid email or password."
    }

    render.Page(c, http.StatusOK, authpages.LoginPage(csrf.GetToken(c), errorMsg))
}

func HandleLogin(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        email := c.PostForm("email")
        password := c.PostForm("password")

        var user models.User
        if err := db.Where("email = ?", email).First(&user).Error; err != nil {
            c.Redirect(http.StatusFound, "/login?error=invalid")
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
            c.Redirect(http.StatusFound, "/login?error=invalid")
            return
        }

        session := sessions.Default(c)
        session.Set("user_id", user.ID.String())
        session.Set("user_name", user.Name)
        session.Set("role", user.Role)
        session.Set("business_id", user.BusinessID.String())
        if user.BranchID != nil {
            session.Set("branch_id", user.BranchID.String())
        }
        session.Save()

        c.Redirect(http.StatusFound, dashboardFor(user.Role))
    }
}

func HandleLogout(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
    c.Redirect(http.StatusFound, "/login")
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