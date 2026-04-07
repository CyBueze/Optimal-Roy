package main

import (
    "net/http"

    "github.com/CyBueze/org/views"
    "github.com/CyBueze/org/internal/config"
    "github.com/CyBueze/org/internal/database"
    "github.com/CyBueze/org/internal/models"
    authhandler "github.com/CyBueze/org/internal/handlers/auth"

    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    csrf "github.com/utrack/gin-csrf"
)

func main() {
  
  cfg := config.Load()
  db := database.Connect(cfg.DatabaseURL)
  db.AutoMigrate(
    &models.Business{},
    &models.Branch{},
    &models.User{},
)
  
	r := gin.Default()
	
	//session store
	store := cookie.NewStore([]byte(cfg.SessionSecret))
    store.Options(sessions.Options{
        Path:     "/",
        MaxAge:   86400 * 7,
        HttpOnly: true,
        Secure:   cfg.AppEnv == "production",
        SameSite: http.SameSiteStrictMode,
    })
    r.Use(sessions.Sessions("org-session", store))
    
    // csrf
    
    r.Use(csrf.Middleware(csrf.Options{
    Secret: cfg.SessionSecret,
    ErrorFunc: func(c *gin.Context) {
        c.String(http.StatusForbidden, "CSRF token mismatch")
        c.Abort()
    },
    }))

// routes

	r.GET("/", func(c *gin.Context) {
		views.InventoryPage().Render(c.Request.Context(), c.Writer)
	})
	
	r.GET("/sales", func(c *gin.Context) {
		views.Calculator().Render(c.Request.Context(), c.Writer)
	})
	
	r.GET("/cart", func(c *gin.Context){
	  views.CartPage().Render(c.Request.Context(), c.Writer)
	})
	
	r.GET("/login", authhandler.ShowLogin)
  r.POST("/login", authhandler.HandleLogin(db))
  r.POST("/logout", authhandler.HandleLogout)

	r.Run(":" + cfg.Port)
}