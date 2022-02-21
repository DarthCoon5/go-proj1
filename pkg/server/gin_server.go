package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//ServerInterface ...
type ServerInterface interface {
	RepositoryStatus() error
	AddRoutes(router *gin.Engine) error
}

type Server struct {
	host            string
	serverInterface ServerInterface
}

func NewServer(newHost string, newServerInterface ServerInterface) *Server {
	return &Server{host: newHost, serverInterface: newServerInterface}
}

func setupCors(r *gin.Engine) {
	r.Use(
		cors.New(cors.Config{
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{
				"Content-Type", "Content-Length", "Accept-Encoding",
			},
			AllowOriginFunc: func(origin string) bool {
				return true
			},
		}),
	)
}

func (s Server) Run() error {
	if s.host != "" {
		gin.SetMode(gin.ReleaseMode)
		router := gin.Default()
		setupCors(router)
		err := s.serverInterface.AddRoutes(router)
		if err != nil {
			return err
		}
		router.GET("/check_alive", func(c *gin.Context) {
			err := s.serverInterface.RepositoryStatus()
			if err != nil {
				c.JSON(http.StatusOK, err.Error())
			}
			c.JSON(http.StatusOK, "OK")
		})
		srv := &http.Server{
			Addr:    s.host,
			Handler: router,
		}
		go func() {
			// service connections
			if err := srv.ListenAndServe(); err != nil {
				log.Printf("listen: %s\n", err)
			}
		}()
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("Shutdown Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			return errors.New("Server Shutdown:" + err.Error())
		}
		return errors.New("Server exiting")
	} else {
		fmt.Println("ERROR: RouteConfigs = nil")
	}
	return errors.New("ERROR: host = ''")
}
