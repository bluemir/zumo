package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) listChannels(c *gin.Context) {
	// TOOD filter by username
	channels, err := server.backend.GetChannels()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, channels)
}
func (server *Server) createChannel(c *gin.Context) {
	// TODO check auth
	req := &struct {
		Name   string
		Labels map[string]string
	}{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	channel, err := server.backend.CreateChannel(req.Name, req.Labels)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, channel)
}
func (server *Server) deleteChannel(c *gin.Context) {
	channelID := c.Param("channelID")
	err := server.backend.DeleteChannel(channelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
func (server *Server) joinChannel(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.MustGet(keyUsername).(string)

	err := server.backend.Join(channelID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
func (server *Server) leaveChannel(c *gin.Context) {

}
func (server *Server) invite(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.Param("username")

	err := server.backend.Join(channelID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func (server *Server) kick(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.Param("username")

	err := server.backend.Leave(channelID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
