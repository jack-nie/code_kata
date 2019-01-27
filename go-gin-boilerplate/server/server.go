package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/controllers"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/middlewares"
	"github.com/sirupsen/logrus"
)

func Init() {
	r := NewRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Infof("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Infof("Server Shutdown: %s", err)
	}
	logrus.Info("Server exiting")
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			user := new(controllers.UsersController)
			userGroup.POST("/signup", user.Signup)
			userGroup.GET("/:id", user.Show)
		}
	}

	admin := router.Group("/admin")
	a := new(controllers.AdminController)
	admin.POST("/", a.Create)
	admin.POST("/login", a.Login)
	admin.Use(middlewares.Auth("admin"))
	{
		post := new(controllers.PostsController)
		admin.POST("/posts/", post.Create)
	}
	return router
}
