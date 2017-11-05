package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createBot(c *gin.Context) {
	req := &struct {
		Name   string
		Driver string
	}{}

	c.Bind(req)

	if err := server.pod.CreateBot(req.Name, req.Driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
