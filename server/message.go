package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (server *Server) findMessages(c *gin.Context) {
	// TODO
	channelID := c.Param("channelID")

	msgs, err := server.backend.GetMessages(channelID, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, msgs)
}

func (server *Server) postMessage(c *gin.Context) {
	channelID := c.Param("channelID")
	username := c.MustGet(keyUsername).(string)
	msg := &struct {
		Text string
	}{}

	if err := c.Bind(msg); err != nil {
		logrus.Warn(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	server.backend.AppendMessage(username, channelID, msg.Text, []byte("{}"))
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
