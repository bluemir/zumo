package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createHook(c *gin.Context) {
	req := &struct {
		ChannelID string
		Username  string
	}{}

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	hook, err := server.backend.CreateHook(req.ChannelID, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hook)
}

func (server *Server) doHook(c *gin.Context) {
	hookID := c.Param("hookID")
	req := &struct {
		Text   string
		Detail json.RawMessage
	}{}

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	msg, err := server.backend.DoHook(hookID, req.Text, req.Detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}
