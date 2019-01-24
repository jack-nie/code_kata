package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() {
	config = config.GetConfig()
	r := NewRouter()
	srv := &http.server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown:", err)
	}
	logrus.Println("Server exiting")
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//health := new(controllers.HealthController)

	//router.GET("/health", health.Status)
	//router.Use(middlewares.AuthMiddleware())

	//v1 := router.Group("v1")
	//{
	//	userGroup := v1.Group("user")
	//	{
	//		user := new(controllers.UserController)
	//		userGroup.GET("/:id", user.Retrieve)
	//	}
	//}
	return router
}
