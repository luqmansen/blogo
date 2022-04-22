package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luqmansen/blogo/pkg/config"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"net/http"
)

func New(config *config.Configuration, httpHandler ServerInterface) *http.Server {

	router := gin.New()

	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())
	router = RegisterHandlers(router, httpHandler)

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
	}
}
