package server

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"

	"github.com/bluemir/zumo/backend"
	"github.com/bluemir/zumo/pod"
)

const (
	keyUsername = "USERNAME"
)

func Run(b backend.Backend, p pod.Pod) error {
	app := gin.Default()

	server := &Server{b, p, rice.MustFindBox("../dist")}

	app.StaticFS("/static", server.dist.HTTPBox())

	app.GET("/", server.CheckAuth, func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", server.dist.MustBytes("html/index.html"))
	})

	app.GET("/register", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", server.dist.MustBytes("html/register.html"))
	})
	app.POST("/register", func(c *gin.Context) {
		req := &struct {
			Id       string
			Password string
		}{}
		err := c.Bind(req)
		if err != nil {
			return
		}
		if _, err := server.backend.CreateUser(req.Id, map[string]string{
			"zumo.type": "user",
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if _, err := server.backend.CreateToken(req.Id, req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	app.GET("/ws", server.CheckAuth, server.ws)

	app.GET("/api/v1/users/:username/joinned-channel", server.CheckAuth, server.joinnedChannel)

	app.GET("/api/v1/channels", server.CheckAuth, server.listChannels)
	app.POST("/api/v1/channels", server.CheckAuth, server.createChannel)
	app.DELETE("/api/v1/channels/:channelID", server.CheckAuth, server.deleteChannel)

	app.PUT("/api/v1/channels/:channelID/join", server.CheckAuth, server.joinChannel)
	app.PUT("/api/v1/channels/:channelID/invite/:username", server.CheckAuth, server.invite)

	app.GET("/api/v1/channels/:channelID/messages", server.CheckAuth, server.findMessages)
	app.POST("/api/v1/channels/:channelID/messages", server.CheckAuth, server.postMessage)

	app.POST("/api/v1/bots", server.CheckAuth, server.createBot)

	app.POST("/hooks/:hookID", server.CheckAuth) // hook

	app.GET("/adaptor/geo/:hookID", server.CheckAuth)

	// will ?
	/*app.Any("/plugin/todo", server.CheckAuth, func(c *gin.Context) {

	})
	app.Any("/plugin/geo", server.CheckAuth, func(c *gin.Context) {

	})
	app.Any("/plugin/check-in", func(c *gin.Context) {

	})*/
	return app.Run("localhost:4000")
}

type Server struct {
	backend backend.Backend
	pod     pod.Pod // use kv leadership
	dist    *rice.Box
}

func (server *Server) CheckAuth(c *gin.Context) {
	str := c.GetHeader("Authorization")
	if str == "" {
		// Request Auth
		c.Header("WWW-Authenticate", "Basic realm=Auth required!")
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	token, err := server.backend.Token(str)
	if err != nil {
		c.String(http.StatusUnauthorized, "Token Not Found: %s", err.Error())
		c.Abort()
		return
	}
	c.Set(keyUsername, token.Username)

	return // continue to next
}
