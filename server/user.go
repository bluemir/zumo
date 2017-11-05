package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) joinnedChannel(c *gin.Context) {
	username := c.Param("username")
	if username == "me" {
		username = c.MustGet(keyUsername).(string)
	}
	channelIDs, err := server.backend.JoinnedChannel(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, channelIDs)
		return
	}
	c.JSON(http.StatusOK, channelIDs)
}
