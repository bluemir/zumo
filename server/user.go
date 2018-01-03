package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) register(c *gin.Context) {
	req := &struct {
		ID       string
		Password string
	}{}
	err := c.Bind(req)
	if err != nil {
		return
	}
	if _, err := server.backend.CreateUser(req.ID, map[string]string{
		"zumo.type": "user",
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	if _, err := server.backend.CreateToken(req.ID, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}
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
func (server *Server) getUserInfo(c *gin.Context) {
	username := c.Param("username")
	if username == "me" {
		username = c.MustGet(keyUsername).(string)
	}
	user, err := server.backend.GetUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, user)
		return
	}
	c.JSON(http.StatusOK, user)

}
